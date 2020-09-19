package workers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Worker struct {
	db *mongo.Database
}

func New() *Worker {

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

	return &Worker{db: db}
}

func (w *Worker) DropTest() error {
	ctx := context.Background()
	if err := w.db.Collection("mctest").Drop(ctx); err != nil {
		return err
	}
	if err := w.db.Collection("uctest").Drop(ctx); err != nil {
		return err
	}
	if err := w.db.Collection("octest").Drop(ctx); err != nil {
		return err
	}
	if err := w.db.Collection("rctest").Drop(ctx); err != nil {
		return err
	}
	return nil
}

// Function to generally add anything to any collection
func (w *Worker) Add(collectionName string, i interface{}) (id string, err error) {

	c := w.db.Collection(collectionName)

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
func (w *Worker) DeleteOne(collectionName string, i interface{}) (n int64, err error) {
	c := w.db.Collection(collectionName)
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

// Get gets details from db with given filter
func (w *Worker) Get(collectionName string, i interface{}) (l []interface{}, err error) {
	c := w.db.Collection(collectionName)
	ctx := context.Background()
	typ := reflect.TypeOf(i)
	a := reflect.Zero(reflect.TypeOf(i)).Interface()

	// log.Println("Without Filter : ", reflect.DeepEqual(a, i))

	if reflect.DeepEqual(a, i) {
		i = bson.D{}
	}

	cur, err := c.Find(ctx, i)

	for cur.Next(ctx) {

		i := reflect.New(typ).Interface()

		// Remember dont use a pointer to l here by i
		if err = cur.Decode(i); err != nil {
			return l, err
		}

		l = append(l, i)

	}
	defer cur.Close(ctx)

	return l, err
}

// GetOne gets single results from db with given filter
func (w *Worker) GetOne(collectionName string, i interface{}) (l interface{}, err error) {
	c := w.db.Collection(collectionName)
	ctx := context.Background()
	typ := reflect.TypeOf(i)
	l = reflect.New(typ).Interface()

	if i == reflect.Zero(reflect.TypeOf(i)).Interface() {
		i = bson.D{}
	}

	r := c.FindOne(ctx, i)

	if r.Err() != nil {
		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {
		return nil, err
	}

	return l, err
}

// FindOneAndUpdate FindOneAndUpdate
func (w *Worker) FindOneAndUpdate(collectionName string, i interface{}, u interface{}) (l interface{}, err error) {

	c := w.db.Collection(collectionName)
	ctx := context.Background()
	typ := reflect.TypeOf(i)

	l = reflect.New(typ).Interface()

	upsert := true
	after := options.After
	r := c.FindOneAndUpdate(ctx, i, bson.M{"$set": u}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	})

	if r.Err() != nil {
		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {
		return nil, err
	}

	return l, err
}

// FindOne FindOne
func (w *Worker) FindOne(collectionName string, i interface{}) (l interface{}, err error) {
	c := w.db.Collection(collectionName)
	ctx := context.Background()
	typ := reflect.TypeOf(i)

	l = reflect.New(typ).Interface()

	r := c.FindOne(ctx, i)

	if err := r.Err(); err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, errors.New("NIL")
		}
		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {
		return nil, err
	}

	return l, err
}

// FindOneWithOr FindOneWithOr
func (w *Worker) FindOneWithOr(collectionName string, i ...interface{}) (l interface{}, err error) {

	c := w.db.Collection(collectionName)
	ctx := context.Background()
	typ := reflect.TypeOf(i[0])

	l = reflect.New(typ).Interface()

	filters := []interface{}{}
	filters = append(filters, i...)

	r := c.FindOne(ctx, bson.D{{Key: "$or", Value: filters}})

	if err := r.Err(); err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, errors.New("NIL")
		}
		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {

		return nil, err
	}

	return l, err
}

// PrintModel PrintModel
func PrintModel(u interface{}) string {
	log.Println("\n________Print Model_______")
	log.Println()
	b, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	s := string(b)
	return s
}
