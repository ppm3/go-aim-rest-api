package configs

import (
	"context"
	"os"
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

	config := &ServerConfig{
		ProjectName:    os.Getenv("PROJECT_NAME"),
		ProjectVersion: os.Getenv("PROJECT_VERSION"),
		Environment:    os.Getenv("ENVIRONMENT"),
		Server: ServerInitConfig{
			Url:  os.Getenv("SERVER_URL"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Mysql: MysqlDBConfig{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Database: os.Getenv("MYSQL_DATABASE"),
		},
		Mongo: MongoDBConfig{
			Host:     os.Getenv("MONGODB_HOST"),
			Port:     os.Getenv("MONGODB_PORT"),
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
			Database: os.Getenv("MONGODB_DATABASE"),
		},
	}

	return config, nil
}
