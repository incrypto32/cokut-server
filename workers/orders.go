package workers

import (
	"context"
	"log"

	"github.com/incrypt0/cokut-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewOrderAggregationHelper() *orderAggregationHelper {
	idConversionStage := bson.D{
		{Key: "$addFields", Value: bson.M{"roid": bson.M{"$toObjectId": "$rid"}}},
	}

	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}

	sortStage := bson.D{{Key: "$sort", Value: bson.M{"time": -1}}}

	userLookupStage := bson.D{
		{
			Key: "$lookup", Value: bson.M{
				"from":         "users",
				"localField":   "uid",
				"foreignField": "uid",
				"as":           "user",
			},
		},
	}

	// pipeline := []bson.M{
	// 	{
	// 		"$match": bson.M{
	// 			"$expr": bson.M{
	// 				"$in": []interface{}{
	// 					bson.M{"$toString": "$_id"},
	// 					"$$meals.k",
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// mealsLookupStage := bson.D{
	// 	{
	// 		Key: "$lookup", Value: bson.M{
	// 			"from": "meals",
	// 			"let": bson.M{
	// 				"meals": bson.M{"$objectToArray": "$items"},
	// 			},
	// 			"pipeline": pipeline,
	// 			"as":       "meals",
	// 		},
	// 	},
	// }

	// restaurantLookupStage := bson.D{
	// 	{
	// 		Key: "$lookup", Value: bson.M{
	// 			"from":         "restaurants",
	// 			"localField":   "roid",
	// 			"foreignField": "_id",
	// 			"as":           "restaurant",
	// 		},
	// 	},
	// }

	// restaurantUnwindStage := bson.D{
	// 	{
	// 		Key: "$unwind", Value: bson.D{
	// 			{Key: "path", Value: "$restaurant"},
	// 			{Key: "preserveNullAndEmptyArrays", Value: false},
	// 		},
	// 	},
	// }
	userUnwindStage := bson.D{
		{
			Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$user"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			},
		},
	}

	return &orderAggregationHelper{
		idConversionStage: idConversionStage,
		sortStage:         sortStage,
		matchStage:        matchStage,
		// mealsLookupStage:      mealsLookupStage,
		// restaurantLookupStage: restaurantLookupStage,
		// restaurantUnwindStage: restaurantUnwindStage,
		userLookupStage: userLookupStage,
		userUnwindStage: userUnwindStage,
	}
}

// PaginateOrders
func (w *Worker) PaginatedOrders(collectionName string, limit int, page int) (l []models.Order, err error) {
	ctx := context.Background()

	if err != nil {
		return nil, err
	}

	var showsLoadedStruct []models.Order

	showLoadedStructCursor, err := w.aggregateOrders(collectionName, limit, page)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	if err = showLoadedStructCursor.All(ctx, &showsLoadedStruct); err != nil {
		log.Println(err)

		return nil, err
	}

	return showsLoadedStruct, err
}

func (w *Worker) aggregateOrders(collectionName string, limit int, page int) (*mongo.Cursor, error) {
	c := w.db.Collection(collectionName)
	ctx := context.Background()

	showLoadedStructCursor, err := c.Aggregate(
		ctx,
		mongo.Pipeline{
			w.oh.idConversionStage,
			w.oh.matchStage,
			w.oh.sortStage,
			w.oh.userLookupStage,
			w.oh.userUnwindStage,
			// w.oh.restaurantLookupStage,
			// w.oh.mealsLookupStage,
			// w.oh.restaurantUnwindStage,
			bson.D{
				{Key: "$skip", Value: (page - 1) * limit},
			},
			bson.D{
				{Key: "$limit", Value: limit},
			},
		},
	)

	return showLoadedStructCursor, err
}

// Get gets details from db with given filter
func (w *Worker) GetAllOrders(collectionName string, limit int64, page int64) (result []models.Order, err error) {
	c := w.db.Collection(collectionName)

	options := options.Find()
	options.SetSort(bson.D{{Key: "time", Value: -1}}).SetLimit(limit).SetSkip((page - 1) * limit)

	ctx := context.Background()

	cur, err := c.Find(ctx, bson.D{}, options)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	if err = cur.All(ctx, &result); err != nil {
		log.Println(err)

		return nil, err
	}

	defer cur.Close(ctx)

	return result, err
}

// Get gets details from db with given filter
func (w *Worker) GetOrdersByUser(
	collectionName string,
	limit int64,
	page int64,
	uid string) (result []models.Order, err error) {

	c := w.db.Collection(collectionName)

	options := options.Find()
	options.SetSort(bson.D{{Key: "time", Value: -1}}).SetLimit(limit).SetSkip((page - 1) * limit)

	ctx := context.Background()

	cur, err := c.Find(ctx, bson.D{{Key: "uid", Value: uid}}, options)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	if err = cur.All(ctx, &result); err != nil {
		log.Println(err)

		return nil, err
	}

	defer cur.Close(ctx)

	return result, err
}
