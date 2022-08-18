package common

import (
	"errors"
	"fmt"
	"time"

	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/interfaces"
)

const MaxTry = 1000

func GenerateUniqueUserName(userDao interfaces.IUserDAO, userName string) (string, error) {
	for i := 0; i < MaxTry; i++ {
		timestamp := time.Now().UnixMilli()
		tryUserName := fmt.Sprintf("%s%d", userName, timestamp)
		if _, err := userDao.Find("user_name=?", tryUserName); err != nil && err == dao.ErrRecordNotFound {
			return tryUserName, nil
		}
	}
	return "", errors.New("can not generate unique username")
}
