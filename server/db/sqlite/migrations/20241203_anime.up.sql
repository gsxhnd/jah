CREATE TABLE IF NOT EXISTS "anime"
(
    "id"           INTEGER  NOT NULL UNIQUE,
    "code"         TEXT,
    "title"        TEXT,
    "cover"        TEXT,
    "bgm_id"       TEXT,
    "publish_date" DATETIME,
    "created_at"   DATETIME NOT NULL default CURRENT_TIMESTAMP,
    "updated_at"   DATETIME,
    PRIMARY KEY ("id")
);
CREATE INDEX IF NOT EXISTS "anime_index_0"
    ON "anime" ("id");
