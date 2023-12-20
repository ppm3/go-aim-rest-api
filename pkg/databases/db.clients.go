package databases

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Clients struct {
	MongoDB *mongo.Client
	MySQL   *sql.DB
}

func NewClients(mongo *mongo.Client, mysql *sql.DB) *Clients {
	return &Clients{
		MongoDB: mongo,
		MySQL:   mysql,
	}
}
