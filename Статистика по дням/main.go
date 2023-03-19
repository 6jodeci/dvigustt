/*
Тестировал для себя(просто мой снипит) 
На качество кода можно не смотреть я просто дергал кликхаус для изучения
*/
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"

    _ "github.com/ClickHouse/clickhouse-go/v2"
)

func main() {
    // устанавливаем соединение 
    connect, err := sql.Open("clickhouse", "dsn")
    if err != nil {
        log.Fatal(err)
    }
    defer connect.Close()

   // подготовка запроса (можно вынести)
    query := `
        SELECT 
            date,
            COUNT(*) AS events_count, 
            SUM(has_video) AS shows_count, 
            SUM(event = 'click') AS clicks_count, 
            COUNT(DISTINCT ad_id) AS unique_ads_count, 
            COUNT(DISTINCT campaign_union_id) AS unique_campaigns_count 
        FROM ads_data 
        GROUP BY date`

    // выполнение запроса
    rows, err := connect.QueryContext(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // хранилище для результатов запроса
    var (
        date                  string
        events_count          int
        shows_count           int
        clicks_count          int
        unique_ads_count      int
        unique_campaigns_count int
    )
    for rows.Next() {
        err := rows.Scan(&date, &events_count, &shows_count, &clicks_count, &unique_ads_count, &unique_campaigns_count)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s %d %d %d %d %d\n", date, events_count, shows_count, clicks_count, unique_ads_count, unique_campaigns_count)
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
}
