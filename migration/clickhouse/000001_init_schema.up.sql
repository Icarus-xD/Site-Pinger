CREATE TABLE IF NOT EXISTS endpoints_stats (
    id UUID DEFAULT generateUUIDv4(),
    endpoint String,
    created_at DateTime DEFAULT now()
) ENGINE = MergeTree()
ORDER BY id;