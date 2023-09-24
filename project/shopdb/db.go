package shopdb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type Product struct {
	ID    string  `json:"id" bson:"_id,omitempty"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
}

func InsertProduct(client *mongo.Client, product Product) error {
	collection := client.Database("mydb").Collection("products")
	_, err := collection.InsertOne(context.TODO(), product)
	return err
}
func DeleteProduct(client *mongo.Client, productName string) error {
	collection := client.Database("mydb").Collection("products")
	filter := bson.M{"name": productName}
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
func UpdateProductPrice(client *mongo.Client, productName string, newPrice float64) error {
	collection := client.Database("mydb").Collection("products")
	filter := bson.M{"name": productName}
	update := bson.M{"$set": bson.M{"price": newPrice}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
func FindProductByName(client *mongo.Client, productName string) (*Product, error) {
	var product Product
	collection := client.Database("mydb").Collection("products")
	filter := bson.M{"name": productName}
	err := collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
