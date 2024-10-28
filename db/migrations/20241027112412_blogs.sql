-- +goose Up
CREATE TABLE blogs (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    description TEXT,
    cover_image TEXT,
    readable_publish_date TEXT,
    social_image TEXT,
    tag_list TEXT,
    tags TEXT,
    slug TEXT,
    path TEXT,
    url TEXT,
    canonical_url TEXT,
    comments_count INTEGER DEFAULT 0,
    positive_reactions_count INTEGER DEFAULT 0,
    public_reactions_count INTEGER DEFAULT 0,
    collection_id INTEGER,
    created_at TEXT NOT NULL DEFAULT (datetime('now')),
    edited_at TEXT,
    published_at TEXT,
    last_comment_at TEXT,
    published_timestamp TEXT,
    reading_time_minutes INTEGER,
    username TEXT,
    user_full_name TEXT,
    user_profile_image TEXT,
    user_profile_image_90 TEXT,
    organization_name TEXT,
    organization_username TEXT,
    organization_profile_image TEXT,
    organization_profile_image_90 TEXT,
    organization_slug TEXT,
    type_of TEXT
);

-- +goose Down
DROP TABLE IF EXISTS blogs;
