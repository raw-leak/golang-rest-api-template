package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents the MongoDB client and database connection.
type MongoDB struct {
	Client          *mongo.Client
	Database        *mongo.Database
	UsersCollection *mongo.Collection
}

func (m *MongoDB) Disconnect(ctx context.Context) {
	log.Println("Disconnecting MongoDB...")

	m.Client.Disconnect(ctx)

	// if err := m.Client.Disconnect(ctx); err != nil {
	// 	log.Printf("Failed to disconnect from MongoDB: %v\n", err)
	// }
	log.Println("MongoDB disconnected")
}

// New initializes and returns a MongoDB instance.
func New(ctx context.Context, uri string, dbName string) (*MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &MongoDB{
		Client:          client,
		Database:        db,
		UsersCollection: db.Collection("users"),
	}, nil
}
