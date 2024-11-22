package scraper

import (
	"strings"

	"github.com/ahmadnouh97/blog-scraper/internal/blog"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
)

func ScrapeDevTo(perPage string, page string, sortBy string, sortDirection string) ([]*DevToData, error) {
	// Scrape Dev.to
	params := map[string]string{
		"per_page":       perPage,
		"page":           page,
		"sort_by":        sortBy,
		"sort_direction": sortDirection,
	}

	return FetchBlogs(params)
}

func ScrapeBlogs(blogRepo *blog.Repository, logger *utils.CustomLogger) {
	// Scrape Dev.to
	devToBlogs, err := ScrapeDevTo("60", "0", "published_at", "desc")

	if err != nil {
		logger.Error("Failed to scrape Dev.to: %v", err)
		return
	}

	// Save blogs to database
	savedBlogs := 0
	for _, devToBlog := range devToBlogs {
		newBlog := &blog.Blog{
			ID:                         devToBlog.ID,
			Title:                      devToBlog.Title,
			Description:                devToBlog.Description,
			CoverImage:                 devToBlog.CoverImage,
			ReadablePublishDate:        devToBlog.ReadablePublishDate,
			SocialImage:                devToBlog.SocialImage,
			TagList:                    strings.Join(devToBlog.TagList, ","),
			Tags:                       devToBlog.Tags,
			Slug:                       devToBlog.Slug,
			Path:                       devToBlog.Path,
			URL:                        devToBlog.URL,
			CanonicalURL:               devToBlog.CanonicalURL,
			CommentsCount:              devToBlog.CommentsCount,
			PositiveReactionsCount:     devToBlog.PositiveReactionsCount,
			PublicReactionsCount:       devToBlog.PublicReactionsCount,
			CollectionID:               devToBlog.CollectionID,
			CreatedAt:                  devToBlog.CreatedAt,
			EditedAt:                   devToBlog.EditedAt,
			PublishedAt:                devToBlog.PublishedAt,
			LastCommentAt:              devToBlog.LastCommentAt,
			PublishedTimestamp:         devToBlog.PublishedTimestamp,
			ReadingTimeMinutes:         devToBlog.ReadingTimeMinutes,
			Username:                   devToBlog.User.Username,
			UserFullName:               devToBlog.User.Name,
			UserProfileImage:           devToBlog.User.ProfileImage,
			UserProfileImage90:         devToBlog.User.ProfileImage90,
			OrganizationName:           devToBlog.Organization.Name,
			OrganizationUsername:       devToBlog.Organization.Username,
			OrganizationProfileImage:   devToBlog.Organization.ProfileImage,
			OrganizationProfileImage90: devToBlog.Organization.ProfileImage90,
			OrganizationSlug:           devToBlog.Organization.Slug,
			TypeOf:                     devToBlog.TypeOf,
		}

		id, err := blogRepo.AddBlog(newBlog)
		if err != nil {
			logger.Error("Failed to save a blog to database: %v", err)
			// TODO: save errors to a file to track errors
			continue
		}

		if id == -1 {
			logger.Warning("Blog with ID %d already exists", newBlog.ID)
			continue
		}

		savedBlogs++
	}

	logger.Info("%v blogs saved to database", savedBlogs)
}
