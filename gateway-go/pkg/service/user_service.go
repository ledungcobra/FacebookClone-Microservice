package service

import (
	"github.com/pkg/errors"
	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/controllers/users/request"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
	"log"
)

type UserService struct {
	userDao interfaces.IUserDAO
}

func NewUserService(userDao interfaces.IUserDAO) *UserService {
	return &UserService{userDao: userDao}
}

func (u *UserService) FindByEmail(email string) (*models.User, error) {
	return u.userDao.Find("email=?", email)
}

func (u *UserService) Save(user *models.User) error {
	return u.userDao.Save(user)
}

func (u *UserService) Verify(email string, token string) error {
	user, err := u.userDao.Find("email=?", email)
	if err != nil {
		if err == dao.ErrRecordNotFound {
			return ErrRecordNotfound
		}
		return err
	}
	if user.Verified {
		return ErrUserAlreadyVerified
	}
	if user.VerificationToken != token {
		return ErrInvalidToken
	}
	user.Verified = true
	claim, err := common.ExtractFromString(token)
	if err != nil {
		return ErrInvalidToken
	}
	if err := claim.Valid(); err != nil {
		return ErrInvalidToken
	}
	if err := u.userDao.Save(user); err != nil {
		return err
	}
	return u.userDao.Save(user)
}

func (u *UserService) ResetPassword(email string) (string, error) {
	user, err := u.FindByEmail(email)
	if err != nil {
		return "", err
	}
	err = u.userDao.RawQuery("delete from codes where user_id = ?", user.ID)
	if err != nil {
		return "", err
	}
	code := models.Code{
		Code:   common.GenerateCode(5),
		UserID: user.ID,
	}
	user.Code = code
	if err := u.Save(user); err != nil {
		return "", err
	}
	return code.Code, nil
}

func (u *UserService) Register(registerRequest request.RegisterRequest) (*models.User, error) {
	var err error
	user := mapRegisterRequestToUser(registerRequest)
	if user.Password, err = common.HashPassword(registerRequest.Password); err != nil {
		return nil, ErrHashingPassword
	}
	userName := registerRequest.FirstName + registerRequest.LastName
	if user.UserName, err = common.GenerateUniqueUserName(u.userDao, userName); err != nil {
		log.Println("Error when generating username ", err)
		return nil, ErrGenerateUniqueName
	}

	if err := u.userDao.Save(&user); err != nil {
		if _, e := u.userDao.Find("email=?", user.Email); e == nil {
			return nil, ErrDuplicateEmail
		}
		return nil, errors.Wrap(err, "")
	}
	return &user, nil
}

func (u *UserService) FindByID(id uint) (*models.User, error) {
	return u.userDao.Find("id=?", id)
}
