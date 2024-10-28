package blog

import (
	"database/sql"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) AddBlog(blog *Blog) (int64, error) {
	query := `
		INSERT INTO blogs (
			id, title, content, description, cover_image, readable_publish_date, social_image, tag_list, tags, slug, 
			path, url, canonical_url, comments_count, positive_reactions_count, public_reactions_count, collection_id, 
			created_at, edited_at, published_at, last_comment_at, published_timestamp, reading_time_minutes, username, 
			user_full_name, user_profile_image, user_profile_image_90, organization_name, organization_username, 
			organization_profile_image, organization_profile_image_90, organization_slug, type_of
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(query, blog.ID, blog.Title, blog.Content, blog.Description, blog.CoverImage,
		blog.ReadablePublishDate, blog.SocialImage, blog.TagList, blog.Tags, blog.Slug, blog.Path, blog.URL,
		blog.CanonicalURL, blog.CommentsCount, blog.PositiveReactionsCount, blog.PublicReactionsCount,
		blog.CollectionID, blog.CreatedAt, blog.EditedAt, blog.PublishedAt, blog.LastCommentAt,
		blog.PublishedTimestamp, blog.ReadingTimeMinutes, blog.Username, blog.UserFullName,
		blog.UserProfileImage, blog.UserProfileImage90, blog.OrganizationName, blog.OrganizationUsername,
		blog.OrganizationProfileImage, blog.OrganizationProfileImage90, blog.OrganizationSlug, blog.TypeOf,
	)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *Repository) GetBlogs() ([]*Blog, error) {

	query := `
		SELECT id, title, content, description, cover_image, readable_publish_date, social_image, tag_list, tags, slug, 
		path, url, canonical_url, comments_count, positive_reactions_count, public_reactions_count, collection_id, 
		created_at, edited_at, published_at, last_comment_at, published_timestamp, reading_time_minutes, username, 
		user_full_name, user_profile_image, user_profile_image_90, organization_name, organization_username, 
		organization_profile_image, organization_profile_image_90, organization_slug, type_of
		FROM blogs
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	blogs := []*Blog{}
	var createdAt string
	var editedAt string
	var publishedAt string
	var lastCommentAt string
	var publishedTimestamp string

	for rows.Next() {
		var blog Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Description, &blog.CoverImage,
			&blog.ReadablePublishDate, &blog.SocialImage, &blog.TagList, &blog.Tags, &blog.Slug, &blog.Path,
			&blog.URL, &blog.CanonicalURL, &blog.CommentsCount, &blog.PositiveReactionsCount,
			&blog.PublicReactionsCount, &blog.CollectionID, &createdAt, &editedAt, &publishedAt,
			&lastCommentAt, &publishedTimestamp, &blog.ReadingTimeMinutes, &blog.Username,
			&blog.UserFullName, &blog.UserProfileImage, &blog.UserProfileImage90, &blog.OrganizationName,
			&blog.OrganizationUsername, &blog.OrganizationProfileImage, &blog.OrganizationProfileImage90,
			&blog.OrganizationSlug, &blog.TypeOf,
		); err != nil {
			return nil, err
		}

		blog.CreatedAt, _ = time.Parse("2006-01-02 15:04:05-07:00", createdAt)
		blog.EditedAt, _ = time.Parse("2006-01-02 15:04:05-07:00", editedAt)
		blog.PublishedAt, _ = time.Parse("2006-01-02 15:04:05-07:00", publishedAt)
		blog.LastCommentAt, _ = time.Parse("2006-01-02 15:04:05-07:00", lastCommentAt)
		blog.PublishedTimestamp, _ = time.Parse("2006-01-02 15:04:05-07:00", publishedTimestamp)

		blogs = append(blogs, &blog)
	}

	return blogs, nil
}
