package config

import (
	"auth-service/helper"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

var (
	ROOT_PATH    string
	APP_PORT     string
	GRPC_PORT    string
	GRPC_TIMEOUT int

	PSQL_CONNECTION string

	SERVICE_CLIENT_USER_GRPC_ADDRESS    string
	SERVICE_CLIENT_ADAPTER_GRPC_ADDRESS string
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	index := strings.LastIndex(basepath, "/config")
	if index != -1 {
		ROOT_PATH = strings.Replace(basepath, "/config", "", index)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error loading .env file: %v", err)
	}
	APP_PORT = helper.GetENV("APP_PORT", "")
	GRPC_PORT = helper.GetENV("GRPC_PORT", "")
	PSQL_CONNECTION = helper.GetENV("PSQL_CONNECTION", "")
	SERVICE_CLIENT_USER_GRPC_ADDRESS = helper.GetENV("SERVICE_CLIENT_USER_GRPC_ADDRESS", "")
	SERVICE_CLIENT_ADAPTER_GRPC_ADDRESS = helper.GetENV("SERVICE_CLIENT_ADAPTER_GRPC_ADDRESS", "")
	GRPC_TIMEOUT = cast.ToInt(helper.GetENV("GRPC_TIMEOUT", ""))
}

func GetPath(dir string) string {
	return fmt.Sprintf("%s/%s", ROOT_PATH, strings.Trim(dir, "/"))
}
