package redisAction

import (
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/mock"
)

type MockRedisActions struct {
	mock.Mock
}

func (m *MockRedisActions) Connect() (*redis.Client, error) {
	args := m.Called()
	return args.Get(0).(*redis.Client), args.Error(1)
}

func (m *MockRedisActions) Ping(rc *redis.Client) (bool, error) {
	args := m.Called(rc)
	return args.Bool(0), args.Error(1)
}
