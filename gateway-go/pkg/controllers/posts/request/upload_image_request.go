package request

type UploadImageRequest struct {
	Data []byte `json:"data"`
	Name string `json:"name"`
}
