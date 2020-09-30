package workers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"

	"github.com/incrypt0/cokut-server/brokers/myerrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Worker .
type Worker struct {
	db *mongo.Database
}

// New Worker
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

// DropTest frop
func (w *Worker) DropTest() error {
	ctx := context.Background()

	if err := w.db.Collection("mctest").Drop(ctx); err != nil {
		return err
	}

	if err := w.db.Collection("test").Drop(ctx); err != nil {
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

//Add Function to generally add anything to any collection
func (w *Worker) Add(collectionName string, i interface{}) (id string, err error) {
	ctx := context.Background()
	c := w.db.Collection(collectionName)
	result, err := c.InsertOne(ctx, i)

	if err != nil {
		log.Println(err)
		return id, errors.New("an error occurred please try again")
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()

	return id, err
}

//DeleteOne Function to generally add anything to any collection
func (w *Worker) DeleteOne(collectionName string, i interface{}) (n int64, err error) {
	c := w.db.Collection(collectionName)
	ctx := context.Background()
	result, err := c.DeleteOne(ctx, i)

	if err != nil {
		log.Println(err)
		return n, errors.New("an error occurred please try again")
	}

	if result.DeletedCount == 0 {
		return n, myerrors.ErrNoRecordsDeleted
	}

	n = result.DeletedCount

	return n, err
}

// Get gets details from db with given filter
func (w *Worker) Get(collectionName string, i interface{}) (l []interface{}, errerr error) {
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
			log.Println(err)
			return l, err
		}

		l = append(l, i)
	}

	defer cur.Close(ctx)

	return l, err
}

// FindOneAndUpdate FindOneAndUpdate
func (w *Worker) FindOneAndUpdate(collectionName string, filter interface{}, update interface{}) (
	l interface{},
	err error) {
	return w.findOneAndUpdateHelper(collectionName, filter, update, false, "")
}

// FindOneAndUpdate FindOneAndUpdate
func (w *Worker) FindOneAndPush(collectionName string, filter interface{}, update interface{}, field string) (
	l interface{},
	err error) {
	return w.findOneAndUpdateHelper(collectionName, filter, update, true, field)
}

// FindOneAndUpdateHelper FindOneAndUpdate
func (w *Worker) findOneAndUpdateHelper(
	collectionName string, i interface{}, u interface{}, push bool, field string) (
	l interface{},
	err error) {
	c := w.db.Collection(collectionName)
	ctx := context.Background()
	filterTyp := reflect.TypeOf(i)
	updateTyp := reflect.TypeOf(u)

	uChecker := reflect.Zero(updateTyp).Interface()

	if reflect.DeepEqual(u, uChecker) {
		log.Println("EMPTY INTERFACE")
		return l, errors.New("EMPTY")
	}

	l = reflect.New(filterTyp).Interface()

	action := "$set"

	if push {
		action = "$push"
		u = bson.M{field: u}
	}

	upsert := true
	after := options.After
	r := c.FindOneAndUpdate(ctx, i, bson.M{action: u}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	})

	if r.Err() != nil {
		log.Println(err)

		if r.Err() == mongo.ErrNoDocuments {
			return nil, myerrors.ErrNIL
		}

		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {
		log.Println(err)

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
			return nil, myerrors.ErrNIL
		}

		log.Println(err)

		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {
		log.Println(err)
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

	for _, r := range i {
		a := reflect.Zero(typ).Interface()
		if !reflect.DeepEqual(a, r) {
			filters = append(filters, r)
		}
	}

	r := c.FindOne(ctx, bson.D{{Key: "$or", Value: filters}})

	if err := r.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, myerrors.ErrNIL
		}

		log.Println(err)

		return nil, err
	}

	// Remember dont use a pointer to l here
	if err = r.Decode(l); err != nil {
		log.Println(err)
		return nil, err
	}

	return l, err
}

// ModelToString            ModelToString
func ModelToString(u interface{}) string {
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
