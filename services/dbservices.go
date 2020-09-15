package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	UserCollection        *mongo.Collection
	RestaurantsCollection *mongo.Collection
	MealsCollection       *mongo.Collection
}

var db *mongo.Database
var ctx context.Context
var C *Collections

func ConnectMongo() *mongo.Database {
	fmt.Println("____Connecting to mongo db____")
	var err error
	var client *mongo.Client
	ctx = context.Background()
	uri := os.Getenv("MONGO_DB_URI")
	opts := options.Client()

	opts.ApplyURI(uri)
	if client, err = mongo.Connect(ctx, opts); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to MongoDB")
	db = client.Database("ecommerce")
	GetCollections()
	return db
}

// Initialize all collections required and save it into the Collections struct
func GetCollections() {
	C = &Collections{
		UserCollection:        db.Collection("users"),
		RestaurantsCollection: db.Collection("restaurants"),
		MealsCollection:       db.Collection("meals"),
	}
}

// Function to generally add anything to any collection
func Add(c *mongo.Collection, i interface{}) (id primitive.ObjectID, err error) {
	fmt.Println("______Add______")
	result, err := c.InsertOne(ctx, i)
	if err != nil {
		return id, err
	}
	//  .........
	if result.InsertedID == nil {
		err = errors.New("An error occured please try again")
		return id, err
	} else {
		id = result.InsertedID.(primitive.ObjectID)
		fmt.Println("Insert Success")
	}
	return id, err

}

func GetAll(c *mongo.Collection, i interface{}) (cur *mongo.Cursor, err error) {

	cur, err = c.Find(ctx, bson.D{})

	return cur, err
}

func GetAll2(c *mongo.Collection, i interface{}) (l []interface{}, err error) {
	typ := reflect.TypeOf(i)

	cur, err := c.Find(ctx, bson.D{})

	for cur.Next(ctx) {

		i := reflect.New(typ).Interface()

		if err = cur.Decode(i); err != nil {
			fmt.Println(err)

			return l, err
		}

		fmt.Println(i)

		l = append(l, i)

	}
	defer cur.Close(ctx)
	fmt.Println()
	fmt.Println(l[0])
	return l, err
}

// Print a model
func PrintModel(u interface{}) string {
	fmt.Println("________Print Model_______")
	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return err.Error()
	}
	s := string(b)
	return s
}
