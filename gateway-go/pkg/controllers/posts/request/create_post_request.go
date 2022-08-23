package request

type CreatePostRequest struct {
	Type       string   `json:"type"`
	Background string   `json:"background"`
	Text       string   `json:"text"`
	Images     []string `json:"images"`
}
