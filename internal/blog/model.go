package blog

import "time"

type Blog struct {
	ID                         int       `json:"id"`
	Title                      string    `json:"title"`
	Description                string    `json:"description"`
	CoverImage                 string    `json:"cover_image"`
	ReadablePublishDate        string    `json:"readable_publish_date"`
	SocialImage                string    `json:"social_image"`
	TagList                    string    `json:"tag_list"`
	Tags                       string    `json:"tags"`
	Slug                       string    `json:"slug"`
	Path                       string    `json:"path"`
	URL                        string    `json:"url"`
	CanonicalURL               string    `json:"canonical_url"`
	CommentsCount              int       `json:"comments_count"`
	PositiveReactionsCount     int       `json:"positive_reactions_count"`
	PublicReactionsCount       int       `json:"public_reactions_count"`
	CollectionID               int       `json:"collection_id"`
	CreatedAt                  time.Time `json:"created_at"`
	EditedAt                   time.Time `json:"edited_at"`
	PublishedAt                time.Time `json:"published_at"`
	LastCommentAt              time.Time `json:"last_comment_at"`
	PublishedTimestamp         time.Time `json:"published_timestamp"`
	ReadingTimeMinutes         int       `json:"reading_time_minutes"`
	Username                   string    `json:"username"`
	UserFullName               string    `json:"user_full_name"`
	UserProfileImage           string    `json:"user_profile_image"`
	UserProfileImage90         string    `json:"user_profile_image_90"`
	OrganizationName           string    `json:"organization_name"`
	OrganizationUsername       string    `json:"organization_username"`
	OrganizationProfileImage   string    `json:"organization_profile_image"`
	OrganizationProfileImage90 string    `json:"organization_profile_image_90"`
	OrganizationSlug           string    `json:"organization_slug"`
	TypeOf                     string    `json:"type_of"`
}

type BlogsPaginationResponse struct {
	Blogs      []*Blog `json:"blogs"`
	Page       int     `json:"page"`
	PageSize   int     `json:"pageSize"`
	TotalItems int     `json:"totalItems"`
	TotalPages int     `json:"totalPages"`
	HasMore    bool    `json:"hasMore"`
}
