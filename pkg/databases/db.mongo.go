package databases

import (
	"context"
	"log"
	"ppm3/go-aim-rest-api/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBActionsI interface {
	MongoConnect() (*mongo.Client, error)
	MongoPing(c *mongo.Client) (bool, error)
	MongoPingDB(c *mongo.Client) (bool, error)
}

type MongoDBActions struct {
	ctx    context.Context
	params *configs.MongoDBConfig
}

func NewMongoDBActions(ctx context.Context, p *configs.MongoDBConfig) MongoDBActionsI {
	return &MongoDBActions{
		ctx:    ctx,
		params: p,
	}
}

func (m *MongoDBActions) MongoConnect() (*mongo.Client, error) {
	// Set connection options
	var mc configs.MongoDBConfig = *m.params
	var uri string

	// Check if port is empty
	if mc.Port != "" {
		uri = mc.Protocol + "://" + mc.Host + ":" + mc.Port
	} else {
		uri = mc.Protocol + "://" + mc.Host
	}

	// Set connection options
	clientOptions := options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(uint64(mc.MaxPoolSize)).
		SetConnectTimeout(time.Duration(mc.ConnectTimeout) * time.Second).
		SetAuth(options.Credential{
			AuthSource: mc.AuthSource,
			Username:   mc.Username,
			Password:   mc.Password,
		})

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(m.ctx, time.Duration(mc.ConnectTimeout)*time.Second)
	defer cancel() // Call the cancel function to avoid a context leak
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Print("[OK] Connected to MongoDB!")

	return client, nil
}

func (m *MongoDBActions) MongoPing(c *mongo.Client) (bool, error) {
	// Check the connection
	err := c.Ping(m.ctx, nil)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (m *MongoDBActions) MongoPingDB(c *mongo.Client) (bool, error) {
	// Check the connection
	var database *mongo.Database = c.Database(m.params.Database)
	err := database.RunCommand(m.ctx, map[string]interface{}{"ping": 1}).Err()
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
