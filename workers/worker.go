package workers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"

	"github.com/incrypt0/cokut-server/brokers/myerrors"
	"github.com/incrypt0/cokut-server/utils"
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
	log.Println(result)

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
func (w *Worker) Get(collectionName string, filter interface{}) (l []interface{}, err error) {
	log.Println(collectionName)
	c := w.db.Collection(collectionName)
	ctx := context.Background()
	typ := reflect.TypeOf(filter)
	a := reflect.Zero(reflect.TypeOf(filter)).Interface()

	// log.Println("Without Filter : ", reflect.DeepEqual(a, i))
	if reflect.DeepEqual(a, filter) {
		filter = bson.D{}
	}

	log.Println(utils.ModelToString(filter))
	log.Println(filter)

	cur, err := c.Find(ctx, filter)

	for cur.Next(ctx) {
		i := reflect.New(typ).Interface()
		// Remember dont use a pointer to l here by i
		if err = cur.Decode(i); err != nil {
			log.Println(err)
			return l, err
		}

		log.Println("HI", i)

		l = append(l, i)
	}

	defer cur.Close(ctx)

	return l, err
}

func (w *Worker) GetMultipleByID(collectionName string, model interface{}, ids []string) (l []interface{}, err error) {
	log.Println(ids, len(ids))
	pids := make([]primitive.ObjectID, len(ids))

	ctx := context.Background()
	c := w.db.Collection(collectionName)
	typ := reflect.TypeOf(model)

	for i, b := range ids {
		log.Println(i, "|", b, ids)
		pids[i], err = primitive.ObjectIDFromHex(b)

		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	cur, err := c.Find(ctx, bson.M{"_id": bson.M{"$in": pids}})
	for cur.Next(ctx) {
		i := reflect.New(typ).Interface()

		// Remember dont use a pointer to l here by i
		if err = cur.Decode(i); err != nil {
			log.Println(err)
			return nil, err
		}

		l = append(l, i)
	}

	return l, err
}

// FindOneAndUpdate FindOneAndUpdate
func (w *Worker) FindOneAndUpdate(collectionName string, filter interface{}, update interface{}) (
	l interface{},
	err error) {
	return w.findOneAndUpdateHelper(collectionName, filter, update, "$set", "")
}

// FindOneAndUpdate FindOneAndUpdate
func (w *Worker) DeleteFromMap(collectionName string, filter interface{}, update interface{}) (
	l interface{},
	err error) {
	return w.findOneAndUpdateHelper(collectionName, filter, update, "$unset", "")
}

// FindOneAndUpdate FindOneAndUpdate
func (w *Worker) FindOneAndPush(collectionName string, filter interface{}, update interface{}, field string) (
	l interface{},
	err error) {
	return w.findOneAndUpdateHelper(collectionName, filter, update, "$push", field)
}

// FindOneAndUpdate FindOneAndUpdate
func (w *Worker) FindOneAndPull(collectionName string, filter interface{}, update interface{}, field string) (
	l interface{},
	err error) {
	return w.findOneAndUpdateHelper(collectionName, filter, update, "$pull", field)
}

// FindOneAndUpdateHelper FindOneAndUpdate
func (w *Worker) findOneAndUpdateHelper(
	collectionName string, i interface{}, u interface{}, action string, field string) (
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

	if action != "$set" && action != "$unset" {
		u = bson.M{field: u}
	}

	upsert := true
	after := options.After
	r := c.FindOneAndUpdate(ctx, i, bson.M{action: u}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	})

	if r.Err() != nil {
		if r.Err() == mongo.ErrNoDocuments {
			return nil, myerrors.ErrNIL
		}

		log.Println(r.Err().Error())

		return nil, r.Err()
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
