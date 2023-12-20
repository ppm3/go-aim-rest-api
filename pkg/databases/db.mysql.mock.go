package databases

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockMySQLActions struct {
	mock.Mock
}

func (m *MockMySQLActions) MySQLConnect() (*sql.DB, error) {
	args := m.Called()
	return args.Get(0).(*sql.DB), args.Error(1)
}

func (m *MockMySQLActions) MySQLPing(c *sql.DB) (bool, error) {
	args := m.Called(c)
	return args.Bool(0), args.Error(1)
}
