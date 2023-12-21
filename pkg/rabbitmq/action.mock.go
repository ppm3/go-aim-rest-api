package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
)

type MockRabbitMQActions struct {
	mock.Mock
}

func (m *MockRabbitMQActions) Connect() (*amqp.Connection, error) {
	args := m.Called()
	return args.Get(0).(*amqp.Connection), args.Error(1)
}

func (m *MockRabbitMQActions) Ping(c *amqp.Connection) (bool, error) {
	args := m.Called(c)
	return args.Bool(0), args.Error(1)
}
