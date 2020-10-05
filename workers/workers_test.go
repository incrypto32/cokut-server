package workers

import (
	"log"
	"testing"
	"time"

	"github.com/incrypt0/cokut-server/utils"
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
	id1, _ := primitive.ObjectIDFromHex("5f689aa84770a61206b9095b")
	id2, _ := primitive.ObjectIDFromHex("5ec0115c23ca01bb29ea922f")

	test1 := Test{
		ID:      id1,
		Name:    "Test 1",
		Message: "Test 1 Success",
		Time:    primitive.Timestamp{T: uint32(time.Now().Unix())},
		Blah:    []Blah{{Hai: "hai", Hmm: "hmm"}},
	}

	test2 := Test{
		ID:      id2,
		Name:    "Test 2",
		Message: "Test 2 Success",
		Time:    primitive.Timestamp{T: uint32(time.Now().Unix())},
	}

	test3 := Test{Message: "Test 1 Edited"}
	w := New()

	if err := w.DropTest(); err != nil {
		t.Error(err)
	}

	c := "test"

	log.Println(test1, test2, test3)

	if i, err := w.Add(c, test1); err != nil {
		t.Log(err)
	} else if i != id1.Hex() {
		t.Error("ERROR")
	}

	if i, err := w.FindOneAndPush(c, Test{ID: id1}, Blah{Hai: "hello",
		Hmm: "HMMMMMM"},
		"blah"); err != nil {
		t.Error(err)
	} else if i.(*Test).Blah[1].Hai != "hello" {
		t.Error(utils.ModelToString(i))
	}

	if l, err := w.GetMultipleByID(c, Test{}, []string{id1.Hex(), id2.Hex()}); err != nil {
		t.Error(err)
	} else {
		t.Log(utils.ModelToString(l))
	}
}
