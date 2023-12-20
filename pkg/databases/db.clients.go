package databases

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Clients struct {
	MongoDB *mongo.Client
	Mysql   *sql.DB
}

func NewClients(m *mongo.Client) *Clients {
	return &Clients{
		MongoDB: m,
	}
}
