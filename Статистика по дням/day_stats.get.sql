SELECT 
    -- выбираем поле date
    date,
    -- считаем количество событий и задаем имя "events_count" для вывода
    COUNT(*) AS events_count, 
    -- считаем количество показов с видео и задаем имя "shows_count" для вывода
    SUM(has_video) AS shows_count, 
    -- считаем количество кликов и задаем имя "clicks_count" для вывода
    SUM(event = 'click') AS clicks_count, 
    -- считаем количество уникальных объявлений и задаем имя "unique_ads_count" для вывода
    COUNT(DISTINCT ad_id) AS unique_ads_count, 
    -- считаем количество уникальных кампаний и задаем имя "unique_campaigns_count" для вывода
    COUNT(DISTINCT campaign_union_id) AS unique_campaigns_count 
    -- выбираем таблицу ads_data
FROM ads_data 
-- группируем по полю date
GROUP BY date 
