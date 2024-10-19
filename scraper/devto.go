package devto

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/ahmadnouh97/blog-scraper/utils"
)

type DevToData struct {
	ClassName                string        `json:"class_name"`
	CloudinaryVideoURL       interface{}   `json:"cloudinary_video_url"`
	ID                       int           `json:"id"`
	Path                     string        `json:"path"`
	PublicReactionsCount     int           `json:"public_reactions_count"`
	ReadablePublishDate      string        `json:"readable_publish_date"`
	ReadingTime              int           `json:"reading_time"`
	Title                    string        `json:"title"`
	Content                  string        `json:"content"`
	UserID                   int           `json:"user_id"`
	PublicReactionCategories []interface{} `json:"public_reaction_categories"`
	CommentsCount            int           `json:"comments_count"`
	VideoDurationString      string        `json:"video_duration_string"`
	PublishedAtInt           int           `json:"published_at_int"`
	TagList                  []interface{} `json:"tag_list"`
	FlareTag                 interface{}   `json:"flare_tag"`
	User                     struct {
		Name           string `json:"name"`
		ProfileImage90 string `json:"profile_image_90"`
		Username       string `json:"username"`
	} `json:"user"`
}

func FetchBlogs(params map[string]string) ([]*DevToData, error) {
	devToData, err := fetchMetadata(params)

	if err != nil {
		return nil, err
	}

	devToData, err = scrapeBlogsContent(devToData)

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

func scrapeBlogsContent(devToData []*DevToData) ([]*DevToData, error) {
	for _, blog := range devToData {
		markdown, err := getBlogMarkdown(blog.Path)

		if err != nil {
			return nil, err
		}
		blog.Content = markdown
	}
	return devToData, nil
}
