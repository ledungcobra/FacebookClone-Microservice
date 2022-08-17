package common

import (
	"errors"
	"fmt"
	"time"

	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
)

const MAX_TRY = 1000

func GenerateUniqueUserName(userDao interfaces.IUserDAO, userName string) (string, error) {
	var user models.User
	for i := 0; i < MAX_TRY; i++ {
		timestamp := time.Now().UnixMilli()
		tryUserName := fmt.Sprintf("%s%d", userName, timestamp)
		if error := userDao.Find(&user, "user_name=?", tryUserName); error != nil && error == dao.ErrRecordNotFound {
			return tryUserName, nil
		}
	}
	return "", errors.New("can not generate unique username")
}
