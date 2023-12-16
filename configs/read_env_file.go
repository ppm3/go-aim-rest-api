package configs

import (
	"context"
	"errors"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type EnvReader interface {
	ReadEnvfile(ctx context.Context, string, projectDirName string) error
}

func ReadEnvfile(ctx context.Context, env string, projectDirName string) error {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	if env == "" {
		return errors.New("environment is not set")
	}

	var envFile string = string(rootPath) + `/.env.` + env

	if envFile != "production" {
		err := godotenv.Load(envFile)
		if err != nil {
			return err
		}
	}

	return nil
}
