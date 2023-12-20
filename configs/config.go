package configs

import (
	"context"
	"os"
	"strconv"
)

type LoadServerConfig struct{}

func NewLoadServerConfig() *LoadServerConfig {
	return &LoadServerConfig{}
}

func (l *LoadServerConfig) Load(ctx context.Context, env string, projectDirName string) (*ServerConfig, error) {

	// Load environment variables
	err := ReadEnvfile(ctx, env, projectDirName)

	if err != nil {
		return nil, err
	}

	var maxPoolSizeStr string = os.Getenv("MONGODB_MAX_POOL_SIZE")
	MaxPoolSize, err := strconv.Atoi(maxPoolSizeStr)
	if err != nil {
		return nil, err
	}
	var ConnectTimeoutStr string = os.Getenv("MONGODB_CONNECT_TIMEOUT")
	ConnectTimeout, err := strconv.Atoi(ConnectTimeoutStr)
	if err != nil {
		return nil, err
	}

	config := &ServerConfig{
		ProjectName:    os.Getenv("PROJECT_NAME"),
		ProjectVersion: os.Getenv("PROJECT_VERSION"),
		Environment:    os.Getenv("ENVIRONMENT"),
		Server: ServerInitConfig{
			Url:  os.Getenv("SERVER_URL"),
			Port: os.Getenv("SERVER_PORT"),
			Mode: os.Getenv("SERVER_MODE"),
		},
		Mysql: MysqlDBConfig{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Database: os.Getenv("MYSQL_DATABASE"),
		},

		Mongo: MongoDBConfig{
			Protocol:       os.Getenv("MONGODB_PROTOCOL"),
			Host:           os.Getenv("MONGODB_HOST"),
			Port:           os.Getenv("MONGODB_PORT"),
			Username:       os.Getenv("MONGODB_USERNAME"),
			Password:       os.Getenv("MONGODB_PASSWORD"),
			Database:       os.Getenv("MONGODB_DATABASE"),
			AuthSource:     os.Getenv("MONGODB_AUTH_SOURCE"),
			MaxPoolSize:    MaxPoolSize,
			ConnectTimeout: ConnectTimeout,
		},
	}

	return config, nil
}
