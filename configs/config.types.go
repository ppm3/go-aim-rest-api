package configs

type ServerConfig struct {
	ProjectName    string
	ProjectVersion string
	Environment    string
	Server         ServerInitConfig
	Mysql          MysqlDBConfig
	Mongo          MongoDBConfig
}

type MysqlDBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type MongoDBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type ServerInitConfig struct {
	Url  string
	Port string
}
