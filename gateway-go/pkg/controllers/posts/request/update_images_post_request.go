package request

type UpdateImagePostRequest struct {
	PostID uint     `json:"post_id"`
	Images []string `json:"images"`
}
