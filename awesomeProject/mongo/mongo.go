package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBConfig represents the MongoDB configuration.
type MongoDBConfig struct {
	URI      string
	Database string
}

// MongoDBClient is a wrapper for the MongoDB client.
type MongoDBClient struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// NewMongoDBClient creates a new MongoDB client with the given configuration.
func NewMongoDBClient(config MongoDBConfig) (*MongoDBClient, error) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")

	// Get a reference to the database and collection
	database := client.Database(config.Database)
	collection := database.Collection("users")

	return &MongoDBClient{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

// CreateUser inserts a new user document into the MongoDB collection.
func (m *MongoDBClient) CreateUser(user interface{}) error {
	_, err := m.collection.InsertOne(context.Background(), user)
	return err
}

// ReadUser retrieves a user document from the MongoDB collection based on a query.
func (m *MongoDBClient) ReadUser(filter bson.M) (interface{}, error) {
	var user interface{}
	err := m.collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

// UpdateUser updates a user document in the MongoDB collection based on a query.
func (m *MongoDBClient) UpdateUser(filter bson.M, update bson.M) error {
	_, err := m.collection.UpdateOne(context.Background(), filter, update)
	return err
}

// DeleteUser deletes a user document from the MongoDB collection based on a query.
func (m *MongoDBClient) DeleteUser(filter bson.M) error {
	_, err := m.collection.DeleteOne(context.Background(), filter)
	return err
}

// Close closes the MongoDB client connection.
func (m *MongoDBClient) Close() error {
	return m.client.Disconnect(context.Background())
}
