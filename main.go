package main

import (
	"auth-service/config"
	"auth-service/server"
)

func getMainServer() *server.Server {
	return &server.Server{
		PSQL_CONNECTION:                     config.PSQL_CONNECTION,
		APP_PORT:                            config.APP_PORT,
		GRPC_PORT:                           config.GRPC_PORT,
		GRPC_TIMEOUT:                        config.GRPC_TIMEOUT,
		SERVICE_CLIENT_USER_GRPC_ADDRESS:    config.SERVICE_CLIENT_USER_GRPC_ADDRESS,
		SERVICE_CLIENT_ADAPTER_GRPC_ADDRESS: config.SERVICE_CLIENT_ADAPTER_GRPC_ADDRESS,
		ROOT_PATH:                           config.ROOT_PATH,
	}
}

func main() {
	s := getMainServer()
	s.Start()
}
