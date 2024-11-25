-- +goose Up

-- 1. Create a New Table Without the content Column
CREATE TABLE blogs_new (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
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
-- 2. Copy Data From the Old Table to the New Table
INSERT INTO blogs_new (id, title, description, cover_image, readable_publish_date, 
    social_image, tag_list, tags, slug, path, url, canonical_url, comments_count, 
    positive_reactions_count, public_reactions_count, collection_id, created_at, 
    edited_at, published_at, last_comment_at, published_timestamp, reading_time_minutes, 
    username, user_full_name, user_profile_image, user_profile_image_90, organization_name, 
    organization_username, organization_profile_image, organization_profile_image_90, 
    organization_slug, type_of)
SELECT id, title, description, cover_image, readable_publish_date, 
    social_image, tag_list, tags, slug, path, url, canonical_url, comments_count, 
    positive_reactions_count, public_reactions_count, collection_id, created_at, 
    edited_at, published_at, last_comment_at, published_timestamp, reading_time_minutes, 
    username, user_full_name, user_profile_image, user_profile_image_90, organization_name, 
    organization_username, organization_profile_image, organization_profile_image_90, 
    organization_slug, type_of
FROM blogs;
-- 3. Drop the Old Table
DROP TABLE blogs;
-- 4. Rename the New Table to the Original Table Name
ALTER TABLE blogs_new RENAME TO blogs;


-- +goose Down
DROP TABLE IF EXISTS blogs;
