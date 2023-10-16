package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"phpToGo/aggregate"
	"time"
)

type MongoRepo struct {
	db       *mongo.Database
	customer *mongo.Collection
}

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

func New(ctx context.Context, connectionString string) (*MongoRepo, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	db := client.Database("demo")
	customers := db.Collection("customers")

	return &MongoRepo{
		db:       db,
		customer: customers,
	}, nil
}

func (m *MongoRepo) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := m.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}
	// Convert to aggregate
	return c.ToAggregate(), nil
}

func (m *MongoRepo) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	_, err := m.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoRepo) Update(customer aggregate.Customer) error {
	//TODO implement me
	panic("implement me")
}
