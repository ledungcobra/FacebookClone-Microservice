package users

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-cmp/cmp"
	"io"
	"io/ioutil"
	"ledungcobra/gateway-go/pkg/app_service"
	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/models"
	"ledungcobra/gateway-go/pkg/service"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestCase struct {
	name          string
	url           string
	method        string
	body          string
	response      common.JSON
	statusCode    int
	expectedError bool
}

func (n *NotificationServiceStub) SendMail(to, subject, body string) (*app_service.SendMailResponse, error) {
	return &app_service.SendMailResponse{
		Success: true,
	}, nil
}

func TestUserController_Register(t *testing.T) {
	userController := getController()
	tests := []TestCase{
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
				"data": common.JSON{
					"success":    true,
					"first_name": "dung",
					"last_name":  "le",
					"verified":   false,
				},
				"errors":  nil,
				"message": "Register user success please active your email to start",
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
		{
			name:   "TestRegister should fail because of validation form",
			url:    "/api/v1/users/register",
			method: "POST",
			body: `{
							"first_name": "dung",
							"last_name": "le",
							"email": "@gmail.com",
							"password":"12345678",
							"birth_year": 1900,
							"birth_month": 1,
							"birth_day": 1,
							"gender":"male"
						}`,
			statusCode: 400,
		},
		{
			name:       "TestRegister should fail because of invalid form",
			url:        "/api/v1/users/register",
			method:     "POST",
			body:       ``,
			statusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *http.Response
			var done bool
			if result, done = UTestStatus(tt.name, t, tt, userController); done {
				return
			}
			if tt.response != nil {
				resp := bodyToJSON(result.Body)
				if resp == nil {
					return
				}
				delete(resp["data"].(common.JSON), "token")
				delete(resp["data"].(common.JSON), "id")
				delete(resp["data"].(common.JSON), "user_name")
				delete(resp["data"].(common.JSON), "picture")
				if diff := cmp.Diff(resp, tt.response, cmp.Comparer(compareResponse)); diff != "" {
					t.Errorf("TestRegister() diff = %v", diff)
				}
			}
		})
	}
}

func UTestStatus(name string, t *testing.T, tt TestCase, userController *fiber.App) (*http.Response, bool) {
	req := httptest.NewRequest(tt.method, tt.url, buildBody(tt.body))
	req.Header.Set("Content-Type", "application/json")
	result, err := userController.Test(req, int(time.Second.Milliseconds()))
	if err != nil {
		t.Errorf(name+") error = %v", err)
		return nil, true
	}
	if result.StatusCode != tt.statusCode {
		t.Errorf(name+" statusCode = %v, want %v", result.StatusCode, tt.statusCode)
	}
	return result, false
}

func getController() *fiber.App {
	app := fiber.New()
	userService := service.NewUserService(NewUserDaoStub(), NewCommonDaoStub[models.Code]())
	userController := NewUserController(&NotificationServiceStub{}, userService)
	v1 := app.Group("/api").Group("/v1")
	userController.RegisterUserRouter(v1)
	return app
}

func compareResponse(actual, expected common.JSON) bool {
	if expected == nil {
		return true
	}
	if len(actual) != len(expected) {
		return false
	}
	for k, v := range expected {
		if k == "id" || k == "token" || k == "user_name" {
			continue
		}
		object, isObject := expected[k].(common.JSON)
		if isObject {
			if !compareResponse(actual[k].(common.JSON), object) {
				return false
			}
		} else {
			if actual[k] != v {
				return false
			}
		}
	}
	return true
}

func bodyToJSON(body io.ReadCloser) common.JSON {
	if body == nil {
		return nil
	}
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

func TestUserController_Login(t *testing.T) {
	app := getController()
	tests := []TestCase{
		{
			name:   "TestLogin should success",
			url:    "/api/v1/users/login",
			method: "POST",
			body: fmt.Sprintf(`{
							"email": "%s",
							"password":"%s"
					}`, TestEmail, TestPassword),
			statusCode: 200,
		},
		{
			name:   "TestLogin should fail when provide wrong password",
			url:    "/api/v1/users/login",
			method: "POST",
			body: fmt.Sprintf(`{
							"email": "%s",
							"password":"%s"
					}`, TestEmail, "12345678aa"),
			statusCode: 400,
		},
		{
			name:   "TestLogin should fail when provide wrong email or password",
			url:    "/api/v1/users/login",
			method: "POST",
			body: fmt.Sprintf(`{
							"email": "%s",
							"password":"%s"
					}`, "abd@gmail.com", "12345678aa"),
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, done := UTestStatus(tt.name, t, tt, app)
			if done {
				return
			}
			if resp != nil {
				t.Log(bodyToJSON(resp.Body))
			}
		})
	}
}
