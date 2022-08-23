package request

type PostRequest struct {
	AuthorID   uint     `json:"author_id""'`
	Text       string   `json:"text"`
	Images     []string `json:"images"`
	Type       string   `json:"type"`
	Background string   `json:"background"`
}
