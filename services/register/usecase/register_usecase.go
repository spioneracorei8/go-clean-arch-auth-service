package usecase

import (
	"auth-service/constants"
	"auth-service/helper"
	"auth-service/models"
	"auth-service/services/register"
	"auth-service/services/user"
	"context"
	"fmt"
	"time"
)

type RegisterUsecase struct {
	registerRepo register.RegisterRepository
	userRepo     user.UserRepository
}

func NewRegisterUsImpl(registerRepo register.RegisterRepository, userRepo user.UserRepository) register.RegisterUsecase {
	return &RegisterUsecase{
		registerRepo: registerRepo,
		userRepo:     userRepo,
	}
}

func (u *RegisterUsecase) RegisterUser(ctx context.Context, params map[string]any, source string) error {
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
	userId, err := u.userRepo.RegisterUser(params)
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

	if err := u.registerRepo.CreateAccount(account); err != nil {
		return err
	}

	return nil
}
