package register

import "auth-service/models"

type RegisterRepository interface {
	CreateAccount(account *models.Account) error
}
