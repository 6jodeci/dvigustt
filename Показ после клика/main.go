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
            click.date, 
            click.ad_id, 
            click.client_union_id, 
            click.campaign_union_id,
            view.event AS view_event, 
            click.event AS click_event,
            click.time AS click_time, 
            view.time AS view_time
        FROM ads_data AS click
        JOIN (
            SELECT *
            FROM ads_data
            WHERE event = 'view'
        ) AS view 
        ON click.ad_id = view.ad_id AND click.client_union_id = view.client_union_id AND click.campaign_union_id = view.campaign_union_id
        WHERE click.event = 'click' AND click.time > view.time`

    // выполнение запроса 
    rows, err := connect.QueryContext(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // хранилище для результатов запроса
    var (
        date            string
        adID            string
        clientUnionID   string
        campaignUnionID string
        viewEvent       string
        clickEvent      string
        clickTime       string
        viewTime        string
    )
    for rows.Next() {
        err := rows.Scan(&date, &adID, &clientUnionID, &campaignUnionID, &viewEvent, &clickEvent, &clickTime, &viewTime)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s %s %s %s %s %s %s %s\n", date, adID, clientUnionID, campaignUnionID, viewEvent, clickEvent, clickTime, viewTime)
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
}
