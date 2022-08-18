package users

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/dao"
	"ledungcobra/gateway-go/pkg/interfaces"
	"ledungcobra/gateway-go/pkg/models"
	"ledungcobra/gateway-go/pkg/service"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type UserDaoStub struct {
	interfaces.IUserDAO
	users     map[uint]*models.User
	CurrentID uint
}

func NewUserDaoStub() *UserDaoStub {
	u := &UserDaoStub{
		users: map[uint]*models.User{
			0: {
				Email:  "test@gmail.com",
				Detail: models.Detail{},
				Post:   nil,
				Model: gorm.Model{
					ID: 0,
				},
			},
		},
	}
	u.CurrentID = 1
	return u
}

func (u *UserDaoStub) Save(user *models.User) error {
	for _, v := range u.users {
		if v.UserName == user.UserName {
			return errors.New("user name is already exist")
		} else if v.Email == user.Email {
			return errors.New("email is already exist")
		}
	}
	user.ID = u.CurrentID
	u.users[user.ID] = user
	u.CurrentID++
	return nil
}

func (u *UserDaoStub) Find(query any, args ...any) (*models.User, error) {
	queryStr, ok := query.(string)
	if !ok {
		return nil, errors.New("query is not string")
	}
	switch {
	case strings.Contains(queryStr, "id"):
		return u.users[args[0].(uint)], nil
	case strings.Contains(queryStr, "email"):
		for _, v := range u.users {
			if v.Email == args[0].(string) {
				return v, nil
			}
		}
	case strings.Contains(queryStr, "user_name"):
		for _, v := range u.users {
			if v.UserName == args[0].(string) {
				return v, nil
			}
		}
	}
	return nil, dao.ErrRecordNotFound
}

type NotificationServiceStub struct {
	interfaces.INotificationService
}

func (n *NotificationServiceStub) SendMail(to, subject, body string) (*service.SendMailResponse, error) {
	return &service.SendMailResponse{
		Success: true,
	}, nil
}

func TestUserController(t *testing.T) {
	app := fiber.New()
	userController := NewUserController(NewUserDaoStub(), &NotificationServiceStub{})
	v1 := app.Group("/api").Group("/v1")
	userController.RegisterUserRouter(v1)
	tests := []struct {
		name     string
		url      string
		method   string
		body     string
		response common.JSON

		statusCode    int
		expectedError bool
	}{
		{
			name:   "TestRegister should success",
			url:    "/api/v1/users/register",
			method: "POST",
			body: `{
							"first_name": "dung",
							"last_name": "le",
							"email": "test1@gmail.com",
							"password":"12345678",
							"birth_year": 1900,
							"birth_month": 1,
							"birth_day": 1,
							"gender":"male"
						}`,

			statusCode: 201,
			response: common.JSON{
				"errors":  nil,
				"message": "Create user success",
			},
		},
		{
			name:   "TestRegister should fail",
			url:    "/api/v1/users/register",
			method: "POST",
			body: `{
							"first_name": "dung",
							"last_name": "le",
							"email": "test@gmail.com",
							"password":"12345678",
							"birth_year": 1900,
							"birth_month": 1,
							"birth_day": 1,
							"gender":"male"
						}`,
			statusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, buildBody(tt.body))
			req.Header.Set("Content-Type", "application/json")
			result, err := app.Test(req, int(time.Second.Milliseconds()))
			if err != nil {
				t.Errorf("TestUserController() error = %v", err)
				return
			}
			if result.StatusCode != tt.statusCode {
				t.Errorf("TestRegister() statusCode = %v, want %v", result.StatusCode, tt.statusCode)
			}
			//if diff := cmp.Diff(bodyToJSON(result.Body), tt.response); diff != "" {
			//	t.Errorf("TestRegister() response = %v, want %v", result.Body, tt.response)
			//}
		})
	}
}

func bodyToJSON(body io.ReadCloser) common.JSON {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("Error:", err)
		return nil
	}
	var result = common.JSON{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println("Error unmarshal:", err)
		return nil
	}
	return result
}

func buildBody(body string) io.Reader {
	return strings.NewReader(body)
}
