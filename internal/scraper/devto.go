package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/ahmadnouh97/blog-scraper/internal/utils"
)

type DevToData struct {
	TypeOf      string `json:"type_of"`
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Content                string      `json:"content"`
	CoverImage             string      `json:"cover_image"`
	ReadablePublishDate    string      `json:"readable_publish_date"`
	SocialImage            string      `json:"social_image"`
	TagList                []string    `json:"tag_list"`
	Tags                   string      `json:"tags"`
	Slug                   string      `json:"slug"`
	Path                   string      `json:"path"`
	URL                    string      `json:"url"`
	CanonicalURL           string      `json:"canonical_url"`
	CommentsCount          int         `json:"comments_count"`
	PositiveReactionsCount int         `json:"positive_reactions_count"`
	PublicReactionsCount   int         `json:"public_reactions_count"`
	CollectionID           int         `json:"collection_id"`
	CreatedAt              time.Time   `json:"created_at"`
	EditedAt               time.Time   `json:"edited_at"`
	CrosspostedAt          interface{} `json:"crossposted_at"`
	PublishedAt            time.Time   `json:"published_at"`
	LastCommentAt          time.Time   `json:"last_comment_at"`
	PublishedTimestamp     time.Time   `json:"published_timestamp"`
	ReadingTimeMinutes     int         `json:"reading_time_minutes"`
	User                   struct {
		Name            string `json:"name"`
		Username        string `json:"username"`
		TwitterUsername string `json:"twitter_username"`
		GithubUsername  string `json:"github_username"`
		WebsiteURL      string `json:"website_url"`
		ProfileImage    string `json:"profile_image"`
		ProfileImage90  string `json:"profile_image_90"`
	} `json:"user"`
	Organization struct {
		Name           string `json:"name"`
		Username       string `json:"username"`
		Slug           string `json:"slug"`
		ProfileImage   string `json:"profile_image"`
		ProfileImage90 string `json:"profile_image_90"`
	} `json:"organization"`
}

func FetchBlogs(params map[string]string) ([]*DevToData, error) {
	devToData, err := fetchMetadata(params)

	if err != nil {
		return nil, err
	}

	// devToData, err = scrapeBlogsContent(devToData)

	if err != nil {
		return nil, err
	}

	return devToData, nil
}

func prepareUrl(params map[string]string) string {
	perPage := params["per_page"]
	page := params["page"]
	sortBy := params["sort_by"]
	sortDirection := params["sort_direction"]
	return fmt.Sprintf("https://dev.to/api/articles/latest?per_page=%s&page=%s&sort_by=%s&sort_direction=%s", perPage, page, sortBy, sortDirection)
}

func fetchMetadata(params map[string]string) ([]*DevToData, error) {
	url := prepareUrl(params)

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error in GET request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	var devToData []*DevToData
	err = json.Unmarshal(body, &devToData)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}

	return devToData, nil
}

func getBlogMarkdown(path string) (string, error) {
	url := fmt.Sprintf("https://dev.to%s", path)

	html, error := utils.FetchHTMLPage(url)

	if error != nil {
		return "", error
	}

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(html)
	if err != nil {
		return "", err
	}

	return markdown, nil
}

// func scrapeBlogsContent(devToData []*DevToData) ([]*DevToData, error) {
// 	for _, blog := range devToData {
// 		markdown, err := getBlogMarkdown(blog.Path)

// 		if err != nil {
// 			return nil, err
// 		}
// 		blog.Content = markdown
// 	}
// 	return devToData, nil
// }
