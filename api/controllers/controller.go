package controllers

import (
	"database/sql"
	"dvigus-tt/internal/config"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
)

// @Summary		Обработка входящего запроса.
// @Description	Обработка входящего запроса на основе предоставленных параметров.
// @Tags			Обработчик запрсоов.
// @Accept			json
// @Produce		json
// @Success		200
// @Failure		429
// @Failure		500
// @Router			/incoming-request [get]
func HandleIncomingRequest(c *gin.Context, db *sql.DB, ipCache *cache.Cache, cfg *config.Config) {
	subnetPrefixSize := cfg.NetPrefix
	requestsPerMinute := cfg.Limit
	cooldownTime := cfg.Cooldown

	ip := getIP(c.Request)
	if ip == "" {
		c.String(http.StatusOK, "Static content")
		return
	}

	subnetIP := getSubnetIP(ip, subnetPrefixSize)
	key := fmt.Sprintf("%s:%d", subnetIP, time.Now().Unix()/60)
	countI, found := ipCache.Get(key)
	var count int
	if found {
		count = countI.(int)
	} else {
		count = 1
		ipCache.Set(key, 1, cache.DefaultExpiration)
	}

	if count >= requestsPerMinute {
		c.Header("Retry-After", strconv.Itoa(int(cooldownTime/time.Second)))
		c.AbortWithStatus(http.StatusTooManyRequests)
		_, found := ipCache.Get(key)
		if found {
			ipCache.Delete(key)
		}
		if _, err := db.Exec("DELETE FROM requests WHERE ip LIKE $1", subnetIP+"%"); err != nil {
			log.Println(err)
		}
		return
	}

	newCount, err := ipCache.IncrementInt(key, 1)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if newCount >= requestsPerMinute {
		time.Sleep(cooldownTime)
	}

	_, err = db.Exec("INSERT INTO requests (ip, datetimez) VALUES ($1, $2)", ip, time.Now().UTC())
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, "Dynamic content")
}

// getIP принимает входящий HTTP-запрос и возвращает IP-адрес клиента.
func getIP(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = req.RemoteAddr
	}
	ips := strings.Split(ip, ", ")
	return ips[len(ips)-1]
}

// getSubnetIP получает адрес подсети (функция из опенсурс рэйтлимитера).
func getSubnetIP(ip string, prefixSize int) string {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return ""
	}
	mask := net.CIDRMask(prefixSize, 32)
	subnetIP := ipAddr.Mask(mask)
	return subnetIP.String()
}

// @Summary		Сброс кеша.
// @Description	Сброс кеша на основе предоставленного префикса.
// @Tags			Сброс кеша.
// @Accept			json
// @Produce		json
// @Success		200
// @Failure		400
// @Failure		500
// @Router			/reset-cache/:prefix [delete]
func ResetHandler(c *gin.Context, ipCache *cache.Cache) {
	prefix := c.Param("prefix")
	found := false
	for k := range ipCache.Items() {
		if strings.HasPrefix(k, prefix) {
			found = true
			ipCache.Delete(k)
		}
	}

	if !found {
		c.String(http.StatusOK, "No cache found for prefix: %s", prefix)
		return
	}

	c.String(http.StatusOK, "Cache reset for prefix: %s", prefix)
}
