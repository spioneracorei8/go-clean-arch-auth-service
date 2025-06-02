package repository

import (
	"auth-service/proto/proto_models"
	"auth-service/services/user"
	"context"
	"time"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcUserRepo struct {
	grpcAddr string
	timeout  int
}

func NewGrpcUserRepoImpl(grpcAddr string, timeout int) user.UserRepository {
	return &grpcUserRepo{
		grpcAddr: grpcAddr,
		timeout:  timeout,
	}
}

func (r *grpcUserRepo) 	RegisterUser(params map[string]any) (*uuid.UUID, error) {
	conn, err := grpc.NewClient(r.grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := proto_models.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.timeout*int(time.Second)))
	defer cancel()

	var registerUserObj = new(proto_models.RegisterUserObj)
	registerUserObj.IdCardNumber = params["id_card_number"].(string)
	registerUserObj.TitleNameTh = params["title_name_th"].(string)
	registerUserObj.FirstNameTh = params["first_name_th"].(string)
	registerUserObj.LastNameTh = params["last_name_th"].(string)
	registerUserObj.TitleNameEn = params["title_name_en"].(string)
	registerUserObj.FirstNameEn = params["first_name_en"].(string)
	registerUserObj.LastNameEn = params["last_name_en"].(string)
	registerUserObj.MobilePhone = params["mobile_phone"].(string)
	registerUserObj.OfficePhone = params["office_phone"].(string)
	registerUserObj.Email = params["email"].(string)
	registerUserObj.Bod = params["bod"].(string)
	registerUserObj.Gender = params["gender"].(string)

	var request = &proto_models.UserRequest{
		RegisterUser: registerUserObj,
	}

	response, err := client.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	if response.UserId == "" {
		return nil, nil
	}

	var userId = uuid.FromStringOrNil(response.UserId)

	return &userId, nil
}
