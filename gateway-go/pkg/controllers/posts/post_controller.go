package posts

import (
	"github.com/gofiber/fiber/v2"
	"ledungcobra/gateway-go/pkg/cloudinary"
	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/controllers/base"
	"ledungcobra/gateway-go/pkg/controllers/posts/request"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/middlewares"
	"ledungcobra/gateway-go/pkg/service"
)

type PostsController struct {
	postService        *service.PostService
	userService        *service.UserService
	uploadImageService interfaces.IUploadImageService
	base.BaseController
}

func NewPostsController(postService *service.PostService,
	userService *service.UserService,
	uploadImageService interfaces.IUploadImageService) *PostsController {
	return &PostsController{
		postService:        postService,
		userService:        userService,
		uploadImageService: uploadImageService,
	}
}

func (p *PostsController) RegisterRoutes(router fiber.Router) {
	post := router.Group("/posts")
	post.Post("/", middlewares.Protected, p.CreatePost)
	post.Post("/uploadImages", middlewares.Protected, p.UploadImages)
}

func (p *PostsController) CreatePost(ctx *fiber.Ctx) error {
	var createPostRequest request.CreatePostRequest
	if err := ctx.BodyParser(&createPostRequest); err != nil {
		return p.InvalidFormResponse(ctx, err)
	}
	userId := ctx.Locals("user_id").(uint)
	result, err := p.postService.Save(createPostRequest, userId)
	if err != nil {
		return p.SendServerError(ctx, err)
	}
	return p.SendOK(ctx, common.JSON{
		"post": common.JSON{
			"id":         result.ID,
			"text":       result.Text,
			"images":     result.Images,
			"background": result.Background,
			"type":       result.Type,
			"author": common.JSON{
				"picture":   result.Author.Picture,
				"id":        result.Author.ID,
				"userName":  result.Author.UserName,
				"firstName": result.Author.FirstName,
				"lastName":  result.Author.LastName,
			},
		},
	}, "Create post success")
}

func (p *PostsController) UploadImages(ctx *fiber.Ctx) error {
	formData, err := ctx.MultipartForm()
	if err != nil {
		return p.InvalidFormResponse(ctx, err)
	}
	files, ok := formData.File["file"]
	if !ok {
		return p.SendBadRequest(ctx, "Missing field data")
	}
	if len(files) == 0 {
		return p.SendBadRequest(ctx, "Missing file")
	}

	areImages := checkTypes(files)
	if !areImages {
		return p.SendBadRequest(ctx, "Files are not images")
	}

	validLength := checkLength(files)
	if !validLength {
		return p.SendBadRequest(ctx, "File size exceeded")
	}

	var results []*cloudinary.UploadResult
	for _, file := range files {
		result, err := p.uploadImageService.UploadFromBytes(extractFromFile(file))
		if err != nil {
			return p.SendServerError(ctx, err)
		}
		results = append(results, result)
	}

	if err != nil {
		return p.SendBadRequest(ctx, "Error when upload image "+err.Error())
	}
	return p.SendOK(ctx, common.JSON{
		"results": toImageResponse(results),
	}, "Upload image success")

}
