package service

import (
	"gorm.io/gorm"
	"ledungcobra/gateway-go/pkg/controllers/posts/request"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
)

type PostService struct {
	postDao *dao.CommonDao[models.Post]
	userDao interfaces.IUserDAO
	db      *gorm.DB
}

func (s *PostService) Save(request request.CreatePostRequest, authorId uint) (*models.Post, error) {
	var post = models.Post{
		AuthorID:   authorId,
		Type:       request.Type,
		Text:       request.Text,
		Images:     request.Images,
		Background: request.Background,
	}
	if result := s.postDao.Save(&post); result != nil {
		return nil, result
	}
	author, err := s.userDao.Find("id=?", authorId)
	if err != nil {
		return nil, err
	}
	post.Author = *author
	return &post, nil
}

func (p *PostService) UpdatePost(postId uint, images []string, userId uint) error {
	post, err := p.FindByID(postId, userId)
	if err != nil {
		return err
	}
	post.Images = images
	if err := p.postDao.Save(post); err != nil {
		return err
	}
	return nil
}

func (s *PostService) FindByID(postId uint, authorID uint) (*models.Post, error) {
	var post models.Post
	r := s.db.Where("id = ? and author_id = ?", postId, authorID).First(&post)
	if r.Error != nil {
		return nil, r.Error
	}
	return &post, nil
}

func NewPostService(
	postDao *dao.CommonDao[models.Post],
	userDao interfaces.IUserDAO,
	db *gorm.DB) *PostService {
	return &PostService{postDao: postDao, db: db, userDao: userDao}
}
