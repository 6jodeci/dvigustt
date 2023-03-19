-- Выбираем все столбцы из таблицы ads_data и создаем алиас click
SELECT *
-- Используем JOIN, чтобы объединить таблицу с самой собой (ads_data), используя алиас
FROM ads_data AS click
JOIN (
    SELECT *
    FROM ads_data
    WHERE event = 'view'
) AS view 
-- Условие JOIN на основе совпадения ad_id, client_union_id, campaign_union_id
ON click.ad_id = view.ad_id AND click.client_union_id = view.client_union_id AND click.campaign_union_id = view.campaign_union_id
-- Ограничиваем результаты выборки только кликами (event = 'click') и просмотрами, которые произошли раньше кликов (click.time > view.time)
WHERE click.event = 'click' AND click.time > view.time