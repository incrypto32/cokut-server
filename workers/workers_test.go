package workers

import (
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	ID      primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string              `json:"name,omitempty" bson:"name,omitempty"`
	Message string              `json:"msg,omitempty" bson:"msg,omitempty"`
	Time    primitive.Timestamp `json:"time,omitempty" bson:"time,omitempty"`
	Blah    string              `json:"blah,omitempty" bson:"blah,omitempty"`
}

func TestDBHandler(t *testing.T) {
	id1, _ := primitive.ObjectIDFromHex("5f689aa84770a61206b9095b")
	id2, _ := primitive.ObjectIDFromHex("5ec0115c23ca01bb29ea922f")

	test1 := Test{
		ID:      id1,
		Name:    "Test 1",
		Message: "Test 1 Success",
		Time:    primitive.Timestamp{T: uint32(time.Now().Unix())},
		Blah:    "blah"}

	test2 := Test{
		ID:      id2,
		Name:    "Test 2",
		Message: "Test 2 Success",
		Time:    primitive.Timestamp{T: uint32(time.Now().Unix())},
		Blah:    "blah"}

	test3 := Test{Message: "Test 1 Edited"}
	_ = New()

	_ = "test"

	log.Println(test1, test2, test3)
}
