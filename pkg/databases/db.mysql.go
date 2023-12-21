package databases

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"ppm3/go-aim-rest-api/configs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ConfigParams struct {
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
	}
}

type MySQLActionsI interface {
	MySQLConnect() (*sql.DB, error)
	MySQLPing(c *sql.DB) (bool, error)
}

type MySQLActions struct {
	ctx    context.Context
	params *configs.MysqlDBConfig
}

func NewMySQLActions(ctx context.Context, p *configs.MysqlDBConfig) MySQLActionsI {
	return &MySQLActions{
		ctx:    ctx,
		params: p,
	}
}

func (m *MySQLActions) MySQLConnect() (*sql.DB, error) {

	// Initialize configParams.Mysql here...
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		m.params.Username,
		m.params.Password,
		m.params.Host,
		m.params.Port,
		m.params.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Duration(m.params.MaxLifeTime) * 3)
	db.SetMaxOpenConns(m.params.MaxOpenCon)
	db.SetMaxIdleConns(m.params.MaxidleCon)

	log.Print("[OK] Connected to Mysql!")

	return db, nil
}

func (m *MySQLActions) MySQLPing(c *sql.DB) (bool, error) {
	err := c.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}
