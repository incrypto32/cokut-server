package workers_test

import (
	"log"
	"testing"

	"github.com/incrypt0/cokut-server/utils"
	"github.com/incrypt0/cokut-server/workers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	ID      primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string              `json:"name,omitempty" bson:"name,omitempty"`
	Message string              `json:"msg,omitempty" bson:"msg,omitempty"`
	Time    primitive.Timestamp `json:"time,omitempty" bson:"time,omitempty"`
	Blah    []Blah              `json:"blah,omitempty" bson:"blah,omitempty"`
}

type Blah struct {
	Hmm string `json:"hmm,omitempty" bson:"hmm,omitempty"`
	Hai string `json:"hai,omitempty" bson:"hai,omitempty"`
}

func TestDBHandler(t *testing.T) {
	w := workers.New()

	a, err := w.PaginateOrders("orders")
	if err != nil {
		log.Println(err)
	}

	log.Println(utils.ModelToString(a))
}
