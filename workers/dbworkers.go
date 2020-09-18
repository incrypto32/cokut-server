package workers

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

func ConnectMongo() *mongo.Database {

	var err error
	var client *mongo.Client
	ctx := context.Background()
	uri := os.Getenv("MONGO_DB_URI")
	opts := options.Client()

	opts.ApplyURI(uri)
	if client, err = mongo.Connect(ctx, opts); err != nil {
		log.Fatal(err)
	}

	db := client.Database("cokut")

	return db
}

// Function to generally add anything to any collection
func Add(c *mongo.Collection, i interface{}) (id string, err error) {

	ctx := context.Background()
	result, err := c.InsertOne(ctx, i)
	if err != nil {
		return id, err
	}

	if result.InsertedID == nil {
		err = errors.New("An error occured please try again")
		return id, err
	} else {
		id = result.InsertedID.(primitive.ObjectID).Hex()
	}

	return id, err
}

// Function to generally add anything to any collection
func DeleteOne(c *mongo.Collection, i interface{}) (n int64, err error) {

	ctx := context.Background()
	result, err := c.DeleteOne(ctx, i)
	if err != nil {
		return n, err
	}
	if result.DeletedCount == 0 {
		err = errors.New("No records were deleted")
		return n, err
	} else {

		n = result.DeletedCount
	}

	return n, err
}

func Get(c *mongo.Collection, i interface{}) (l []interface{}, err error) {
	ctx := context.Background()
	typ := reflect.TypeOf(i)
	a := reflect.Zero(reflect.TypeOf(i)).Interface()

	log.Println("Without Filter : ", reflect.DeepEqual(a, i))

	if reflect.DeepEqual(a, i) {
		i = bson.D{}
	}

	cur, err := c.Find(ctx, i)

	for cur.Next(ctx) {

		i := reflect.New(typ).Interface()

		// Remember dont use a pointer to l here by i
		if err = cur.Decode(i); err != nil {
			log.Println(err)

			return l, err
		}

		l = append(l, i)

	}
	defer cur.Close(ctx)

	return l, err
}

func GetOne(c *mongo.Collection, i interface{}) (l interface{}, err error) {
	ctx := context.Background()
	typ := reflect.TypeOf(i)
	l = reflect.New(typ).Interface()

	fmt.Println("Interface hase zero value : ", i == reflect.Zero(reflect.TypeOf(i)).Interface())

	if i == reflect.Zero(reflect.TypeOf(i)).Interface() {
		i = bson.D{}
	}

	r := c.FindOne(ctx, i)

	if r.Err() != nil {
		log.Println(err)
		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {
		return nil, err
	}

	return l, err
}

// Print a model
func PrintModel(u interface{}) string {
	fmt.Println("\n________Print Model_______")
	fmt.Println()
	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return err.Error()
	}
	s := string(b) + "\n"
	return s
}
