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

	// mysql
	var maxOpenConStr string = os.Getenv("MYSQL_MAX_OPEN_CON")
	maxOpenCon, _ := strconv.Atoi(maxOpenConStr)

	var maxLifeTimeStr string = os.Getenv("MYSQL_MAX_LIFE_TIME")
	maxLifeTime, _ := strconv.Atoi(maxLifeTimeStr)

	var maxidleConStr string = os.Getenv("MYSQL_MAX_IDLE_CON")
	maxidleCon, _ := strconv.Atoi(maxidleConStr)

	// mongodb
	var maxPoolSizeStr string = os.Getenv("MONGODB_MAX_POOL_SIZE")
	maxPoolSize, _ := strconv.Atoi(maxPoolSizeStr)

	var connectTimeoutStr string = os.Getenv("MONGODB_CONNECT_TIMEOUT")
	connectTimeout, _ := strconv.Atoi(connectTimeoutStr)

	// redis
	var redisDialTimeoutStr string = os.Getenv("REDIS_DIAL_TIMEOUT")
	redisDialTimeout, _ := strconv.Atoi(redisDialTimeoutStr)

	var redisReadTimeoutStr string = os.Getenv("REDIS_READ_TIMEOUT")
	redisReadTimeout, _ := strconv.Atoi(redisReadTimeoutStr)

	var redisWriteTimeoutStr string = os.Getenv("REDIS_WRITE_TIMEOUT")
	redisWriteTimeout, _ := strconv.Atoi(redisWriteTimeoutStr)

	var redisPoolSizeStr string = os.Getenv("REDIS_POOL_SIZE")
	redisPoolSize, _ := strconv.Atoi(redisPoolSizeStr)

	var redisMinIdleConnsStr string = os.Getenv("REDIS_MIN_IDLE_CONNS")
	redisMinIdleConns, _ := strconv.Atoi(redisMinIdleConnsStr)

	var redisIdleTimeoutStr string = os.Getenv("REDIS_IDLE_TIMEOUT")
	redisIdleTimeout, _ := strconv.Atoi(redisIdleTimeoutStr)

	var redisMaxRetriesStr string = os.Getenv("REDIS_MAX_RETRIES")
	redisMaxRetries, _ := strconv.Atoi(redisMaxRetriesStr)

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
			Host:        os.Getenv("MYSQL_HOST"),
			Port:        os.Getenv("MYSQL_PORT"),
			Username:    os.Getenv("MYSQL_USERNAME"),
			Password:    os.Getenv("MYSQL_PASSWORD"),
			Database:    os.Getenv("MYSQL_DATABASE"),
			MaxOpenCon:  maxOpenCon,
			MaxLifeTime: maxLifeTime,
			MaxidleCon:  maxidleCon,
		},

		Mongo: MongoDBConfig{
			Protocol:       os.Getenv("MONGODB_PROTOCOL"),
			Host:           os.Getenv("MONGODB_HOST"),
			Port:           os.Getenv("MONGODB_PORT"),
			Username:       os.Getenv("MONGODB_USERNAME"),
			Password:       os.Getenv("MONGODB_PASSWORD"),
			Database:       os.Getenv("MONGODB_DATABASE"),
			AuthSource:     os.Getenv("MONGODB_AUTH_SOURCE"),
			MaxPoolSize:    maxPoolSize,
			ConnectTimeout: connectTimeout,
		},

		RabbitMQ: RabbitMQConfig{
			Host:     os.Getenv("RABBITMQ_HOST"),
			Port:     os.Getenv("RABBITMQ_PORT"),
			Username: os.Getenv("RABBITMQ_USERNAME"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
		},

		Redis: RedisConfig{
			Host:         os.Getenv("REDIS_HOST"),
			Port:         os.Getenv("REDIS_PORT"),
			Password:     os.Getenv("REDIS_PASSWORD"),
			Database:     os.Getenv("REDIS_DB"),
			DialTimeout:  redisDialTimeout,
			ReadTimeout:  redisReadTimeout,
			WriteTimeout: redisWriteTimeout,
			PoolSize:     redisPoolSize,
			MinIdleConns: redisMinIdleConns,
			IdleTimeout:  redisIdleTimeout,
			MaxRetries:   redisMaxRetries,
		},
	}

	return config, nil
}
