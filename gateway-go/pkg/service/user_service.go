package service

import (
	"encoding/json"
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
	codeDao interfaces.ICommonDao[models.Code]
}

func NewUserService(userDao interfaces.IUserDAO, codeDao interfaces.ICommonDao[models.Code]) *UserService {
	return &UserService{userDao: userDao, codeDao: codeDao}
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

func (u *UserService) VerifyCode(code string, email string) bool {
	user, err := u.FindByEmail(email)
	if err != nil {
		return false
	}
	dbCode, err := u.codeDao.Find("user_id=?", user.ID)
	if err != nil {
		return false
	}
	return code == dbCode.Code
}

func (u *UserService) ChangePassword(changePasswordRequest request.ChangePasswordRequest) error {
	var err error
	if changePasswordRequest.Password != changePasswordRequest.ConfPassword {
		return ErrPasswordNotMatch
	}
	claim, err := common.ExtractFromString(changePasswordRequest.Token)
	if err != nil {
		log.Println("Cannot extract from string")
		return ErrInvalidToken
	}
	if err := claim.Valid(); err != nil {
		log.Println("Claim invalid")
		return ErrInvalidToken
	}
	var data common.JSON
	if err := json.Unmarshal([]byte(claim.Subject), &data); err != nil {
		return err
	}
	userId, ok := data["user_id"].(float64)
	if !ok {
		log.Println("Not found user_id")
		return ErrInvalidToken
	}

	user, err := u.FindByEmail(changePasswordRequest.Email)
	if user.ID != uint(userId) {
		log.Println("User id and user.Id do not match")
		return ErrInvalidToken
	}
	if user.Password, err = common.HashPassword(changePasswordRequest.Password); err != nil {
		return ErrHashingPassword
	}
	if err := u.userDao.Save(user); err != nil {
		return err
	}
	return nil
}
