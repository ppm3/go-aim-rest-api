package configs

type ServerConfig struct {
	ProjectName    string
	ProjectVersion string
	Environment    string
	Server         ServerInitConfig
	Mysql          MysqlDBConfig
	Mongo          MongoDBConfig
	RabbitMQ       RabbitMQConfig
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type MysqlDBConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	Database    string
	MaxOpenCon  int
	MaxLifeTime int
	MaxidleCon  int
}

type MongoDBConfig struct {
	Host           string
	Port           string
	Username       string
	Password       string
	Database       string
	Protocol       string
	AuthSource     string
	MaxPoolSize    int
	ConnectTimeout int
}

type ServerInitConfig struct {
	Url  string
	Port string
	Mode string
}
