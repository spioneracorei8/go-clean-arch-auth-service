package adapter

import (
	"auth-service/models"
	"auth-service/proto/proto_models"
)

type GrpcAdapterRepository interface {
	SendMail(mail *models.MailForm) (*proto_models.SendMailResponse, error)
}
