package usecase

import (
	"auth-service/constants"
	"auth-service/helper"
	"auth-service/models"
	"auth-service/services/adapter"
	"auth-service/services/register"
	"auth-service/services/user"
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"
)

type RegisterUsecase struct {
	registerRepo register.RegisterRepository
	userRepo     user.UserRepository
	adapterRepo  adapter.GrpcAdapterRepository
	ROOT_PATH    string
}

func NewRegisterUsImpl(registerRepo register.RegisterRepository, userRepo user.UserRepository, adapterRepo adapter.GrpcAdapterRepository, ROOT_PATH string) register.RegisterUsecase {
	return &RegisterUsecase{
		registerRepo: registerRepo,
		userRepo:     userRepo,
		ROOT_PATH:    ROOT_PATH,
		adapterRepo:  adapterRepo,
	}
}

func (r *RegisterUsecase) RegisterUser(ctx context.Context, params map[string]any, source string) error {
	var (
		username string
	)
	switch source {
	case constants.SOURCE_WEB_APPLICATION, constants.SOURCE_MOBILE_APPLICATION:
		username = params["id_card_number"].(string)
	case constants.SOURCE_WEB_MANAGEMENT:
		username = params["email"].(string)
	}
	fmt.Println("Username:", username)
	// check if account already exists here
	// ----------------------------

	// ----------------------------
	userId, err := r.userRepo.RegisterUser(params)
	if err != nil {
		return err
	}

	var now = helper.NewTimestampFromTime(time.Now())
	account := new(models.Account)
	account.GenUUID()
	account.UserId = userId
	account.Username = params["id_card_number"].(string)
	account.PasswordPlainText = params["password"].(string)
	account.BcryptPwd()
	account.WebAccess = constants.MAP_SOURCE_TO_WEB_ACCESS()[source]
	account.Status = constants.ACCOUNT_STATUS_ACTIVE
	account.CreatedBy = userId.String()
	account.UpdatedBy = userId.String()
	account.CreatedAt = now
	account.UpdatedAt = now

	if err := r.registerRepo.CreateAccount(account); err != nil {
		return err
	}

	var mail = new(models.MailForm)
	mail.To = params["email"].(string)
	mail.ToName = params["first_name_th"].(string) + " " + params["last_name_th"].(string)
	mail.Subject = "สวัสดีครับท่าน"
	mail.Body = ""
	body, err := r.signUpSuccessTemplete(mail.ToName)
	if err != nil {
		return err
	}
	mail.Body = body

	if _, err := r.adapterRepo.SendMail(mail); err != nil {
		return err
	}

	return nil
}

func (r *RegisterUsecase) signUpSuccessTemplete(name string) (string, error) {
	type emailTemplete struct {
		LOGO_URL string
		NAME     string
		REDIRECT string
	}
	tmpl := template.Must(template.ParseFiles("assets/email/sign_up_success.html"))
	logoUrl := helper.GetENV("LOGO_URL", "")
	callbackUrl := helper.GetENV("DOMAIN_APPLICATION_URL", "")

	templeteData := &emailTemplete{
		LOGO_URL: logoUrl,
		NAME:     name,
		REDIRECT: callbackUrl,
	}

	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, templeteData); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
