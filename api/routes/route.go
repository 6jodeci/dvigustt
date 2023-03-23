package routes

import (
	"database/sql"
	"net/http"
	"time"

	"dvigus-tt/api/controllers"
	_ "dvigus-tt/docs"
	"dvigus-tt/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	httpSwagger "github.com/swaggo/http-swagger"
)

// ErrorResponse структура возвращенной ошибки
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRouter(db *sql.DB) *gin.Engine {
	// Инициализируем роутер gin
	router := gin.Default()

	router.GET("/swagger/*any", gin.WrapH(httpSwagger.Handler()))
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	ipCache := cache.New(time.Minute, time.Minute)
	router.GET("/incoming-request", func(c *gin.Context) {
		controllers.HandleIncomingRequest(c, db, ipCache, config.GetConfig())
	})

	router.DELETE("/reset-cache/:prefix", func(c *gin.Context) {
		controllers.ResetHandler(c, ipCache)
	})
	return router
}
