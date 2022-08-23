package cloudinary

type UploadResult struct {
	FileName string
	URL      string
	Type     string
	Width    int
	Height   int
}
