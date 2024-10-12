package devto

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

type DevToResponse struct {
	Result []*DevToData `json:"result"`
}

func Scrape(url string) ([]*DevToData, error) {
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

	var devtoResponse DevToResponse
	err = json.Unmarshal(body, &devtoResponse)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}

	return devtoResponse.Result, nil
}
