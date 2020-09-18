package workers

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	Name    string              `json:"name" bson:"name"`
	Message string              `json:"msg" bson:"msg"`
	Time    primitive.Timestamp `json:"time" bson:"time"`
}

func TestDB(t *testing.T) {
	test1 := &Test{Name: "Test 1 ", Message: "Test 1 Success", Time: primitive.Timestamp{T: uint32(time.Now().Unix())}}
	test2 := &Test{Name: "Test 2 ", Message: "Test 2 Success", Time: primitive.Timestamp{T: uint32(time.Now().Unix())}}
	db := ConnectMongo()
	c := db.Collection("test")

	if err := c.Drop(context.Background()); err != nil {
		t.Error(err)
	}

	if id, err := Add(c, test1); err != nil {
		t.Error(err)
	} else {
		t.Log("collection add test success id :", id)
	}

	if id, err := Add(c, test2); err != nil {
		t.Error(err)
	} else {
		t.Log("collection add test success id :", id)
	}

	if l, err := Get(c, Test{}); err != nil {
		t.Error(err)
	} else {
		t.Log("Get All Success : ")
		t.Log(PrintModel(l))
	}

	if l, err := Get(c, test1); err != nil {
		t.Error(err)
	} else {
		t.Log("Get Filtered Test 1")
		t.Log("Get Filtered Success : ")
		t.Log(PrintModel(l))
	}

	if l, err := Get(c, test2); err != nil {
		t.Error(err)
	} else {
		t.Log("Get Filtered Test 2")
		t.Log("Get Filtered Success : ")
		t.Log(PrintModel(l))
	}

	if l, err := GetOne(c, test1); err != nil {
		t.Error(err)
	} else {
		t.Log("Get One Test 1")
		t.Log("Get One Success : ")
		t.Log(PrintModel(l))
	}

	if l, err := GetOne(c, Test{}); err != nil {
		t.Error(err)
	} else {
		t.Log("Get One Test 2")
		t.Log("Get One Success : ")
		t.Log(PrintModel(l))
	}
}
