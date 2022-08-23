package users

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ledungcobra/gateway-go/pkg/app_service"
	"ledungcobra/gateway-go/pkg/common"
	"ledungcobra/gateway-go/pkg/htmltemplates"
	"ledungcobra/gateway-go/pkg/models"
	"log"
	"os"
	"time"
)

func (u *UserController) sendVerification(ctx *fiber.Ctx, user models.User) (string, error) {
	oneMonth := time.Hour * 24 * 30
	emailVerificationToken, err := u.generateToken(ctx, common.JSON{"email": user.Email}, oneMonth)
	user.VerificationToken = emailVerificationToken
	if err := u.service.Save(&user); err != nil {
		return "", u.SendServerError(ctx, err)
	}

	if err != nil {
		return "", err
	}
	u.handleSendEmailResponse(u.notificationService.SendMail(user.Email, "Verification Email",
		htmltemplates.BuildRegistrationTemplate(user.UserName,
			fmt.Sprintf(os.Getenv("GATEWAY_BASE_FRONTEND_URL")+"/v1/user/verification/token=%s&email=%s", emailVerificationToken, user.Email),
		),
	))

	return emailVerificationToken, nil
}

func (u *UserController) generateToken(ctx *fiber.Ctx, data common.JSON, duration time.Duration) (string, error) {
	emailVerificationToken, err := common.GenerateToken(data, duration)
	if err != nil {
		return "", u.SendServerError(ctx, err)
	}
	return emailVerificationToken, nil
}

func (u *UserController) handleSendEmailResponse(emailResponse *app_service.SendMailResponse, err error) {
	if err != nil {
		log.Println("Error when sending email ", err)
	} else {
		if emailResponse.Success {
			log.Println("Send email success")
		} else {
			log.Println("Send email failed")
		}
	}
}
