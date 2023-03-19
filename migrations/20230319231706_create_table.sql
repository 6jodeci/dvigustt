-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ads_data (
    date Date,
    time DateTime64(0),
    event String,
    platform String,
    ad_id UInt32,
    client_union_id UInt32,
    campaign_union_id UInt32,
    ad_cost_type String,
    ad_cost Float32,
    has_video UInt8,
    target_audience_count UInt32
  ) ENGINE = MergeTree() PARTITION BY toYYYYMMDD(date)
ORDER BY
  (date, time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ads_data;
-- +goose StatementEnd
