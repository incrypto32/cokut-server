package workers

import (
	"context"
	"log"

	"github.com/incrypt0/cokut-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PaginateOrders
func (w *Worker) PaginateOrders(collectionName string) (l interface{}, err error) {
	ctx := context.Background()

	if err != nil {
		return nil, err
	}

	var showsLoadedStruct []models.Order

	showLoadedStructCursor, err := w.aggregateOrders(collectionName, 5, 1)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	if err = showLoadedStructCursor.All(ctx, &showsLoadedStruct); err != nil {
		panic(err)
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
			w.oh.userLookupStage,
			w.oh.restaurantLookupStage,
			w.oh.mealsLookupStage,
			w.oh.userUnwindStage,
			w.oh.restaurantUnwindStage,
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
