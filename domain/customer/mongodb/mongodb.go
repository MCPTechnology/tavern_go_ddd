// mongodb is a mongoDB implementation of the Customer Repository
package mongodb

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbRepository struct {
	db       *mongo.Database
	// Customer is used to store customers
	customer *mongo.Collection
}

// mongoCustomer is an internal type that is used to store a CustomerAggregate inside this repository
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, connectionString string) (*MongodbRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	// Find Metabot DB
	db := client.Database("tavern")
	customers := db.Collection("customer")

	return &MongodbRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (mr *MongodbRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	if err := result.Decode(&c); err != nil {
		return aggregate.Customer{}, err
	}
	// Convert to aggregate
	return c.ToAggregate(), nil
}

func (mr *MongodbRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MongodbRepository) Update(c aggregate.Customer) error {
	panic("to implement")
}
