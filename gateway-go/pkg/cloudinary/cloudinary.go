package cloudinary

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"io"
	"log"
	"os"
)

type CloudinaryUploadService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService() *CloudinaryUploadService {
	cld, err := cloudinary.NewFromParams(os.Getenv("GATEWAY_CLOUD_NAME"), os.Getenv("GATEWAY_CLOUD_API_KEY"), os.Getenv("GATEWAY_CLOUD_SECRET"))
	if err != nil {
		log.Panic("Error occur ", err)
	}
	return &CloudinaryUploadService{cld: cld}
}

func (c *CloudinaryUploadService) UploadFromBytes(reader io.Reader) (*UploadResult, error) {
	result, err := c.cld.Upload.Upload(context.TODO(), reader, uploader.UploadParams{})
	if err != nil {
		return nil, err
	}
	return &UploadResult{
		FileName: result.OriginalFilename,
		URL:      result.URL,
		Type:     result.Type,
		Width:    result.Width,
		Height:   result.Height,
	}, nil
}
