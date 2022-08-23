package posts

import (
	"bytes"
	"io"
	"io/ioutil"
	"ledungcobra/gateway-go/pkg/cloudinary"
	"ledungcobra/gateway-go/pkg/controllers/posts/response"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func checkTypes(files []*multipart.FileHeader) bool {
	for _, file := range files {
		contentType := strings.Split(file.Header.Get("Content-Type"), "/")[1]
		if contentType != "jpeg" && contentType != "png" && contentType != "jpg" {
			return false
		}
	}
	return true
}

func extractFromFile(fileHeader *multipart.FileHeader) io.Reader {
	file, err := fileHeader.Open()
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return bytes.NewReader(data)
}

func checkLength(fileHeaders []*multipart.FileHeader) bool {
	var totalSizeBytes int64
	for _, file := range fileHeaders {
		totalSizeBytes += file.Size
	}
	maxMb, _ := strconv.Atoi(os.Getenv("GATEWAY_MAX_SIZE_IMAGE_UPLOAD_MB"))
	log.Println("Total size: ", totalSizeBytes)
	log.Println("Max mb ", maxMb)
	if totalSizeBytes > int64(maxMb*1024*1024) {
		return false
	}
	return true
}

func toImageResponse(images []*cloudinary.UploadResult) []response.ImageResponse {
	var result []response.ImageResponse
	for _, img := range images {
		result = append(result, response.ImageResponse{
			URL:    img.URL,
			Width:  img.Width,
			Height: img.Height,
		})
	}
	return result
}
