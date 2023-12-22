package databases

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockMongoDBActions struct {
	mock.Mock
}

func (m *MockMongoDBActions) Connect() (*mongo.Client, error) {
	args := m.Called()
	return args.Get(0).(*mongo.Client), args.Error(1)
}

func (m *MockMongoDBActions) Ping(c *mongo.Client) (bool, error) {
	args := m.Called(c)
	return args.Bool(0), args.Error(1)
}

func (m *MockMongoDBActions) PingDB(c *mongo.Client) (bool, error) {
	args := m.Called(c)
	return args.Bool(0), args.Error(1)
}
