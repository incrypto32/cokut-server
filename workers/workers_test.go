package workers

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	Name    string              `json:"name,omitempty" bson:"name,omitempty"`
	Message string              `json:"msg,omitempty" bson:"msg,omitempty"`
	Time    primitive.Timestamp `json:"time,omitempty" bson:"time,omitempty"`
	Blah    string              `json:"blah,omitempty" bson:"blah,omitempty"`
}

func TestDBHandler(t *testing.T) {
	test1 := Test{Name: "Test 1", Message: "Test 1 Success", Time: primitive.Timestamp{T: uint32(time.Now().Unix())}, Blah: "blah"}
	test2 := Test{Name: "Test 2", Message: "Test 2 Success", Time: primitive.Timestamp{T: uint32(time.Now().Unix())}, Blah: "blah"}
	test3 := Test{Message: "Test 1 Edited"}
	w := New()

	c := "test"

	if err := w.DropTest(); err != nil {
		t.Error(err)
	}

	if id, err := w.Add(c, test1); err != nil {
		t.Error(err)
	} else {
		t.Log("collection add test success id :", id)
	}

	if id, err := w.Add(c, test2); err != nil {
		t.Error(err)
	} else {
		t.Log("collection add test success id :", id)
	}

	if l, err := w.Get(c, Test{}); err != nil {
		t.Error(err)
	} else {
		t.Log("Get All Success : ")
		t.Log(PrintModel(l))

		if len(l) != 2 {
			t.Error("Get failed")
		}
	}

	if l, err := w.Get(c, test1); err != nil {
		t.Error(err)
	} else {
		t.Log("Get Filtered Test 1")
		t.Log("Get Filtered Success : ")
		t.Log(PrintModel(l))
		if len(l) != 1 || l[0].(*Test).Name != "Test 1" {
			t.Error("GetFiltered failed (ERROR!!!)")
		} else {
			t.Log("Success")
		}
	}

	if l, err := w.Get(c, test2); err != nil {
		t.Error(err)
	} else {
		t.Log("Get Filtered Test 2  :")
		t.Log(PrintModel(l))
		if len(l) != 1 || l[0].(*Test).Name != "Test 2" {
			t.Error("GetFiltered failed (ERROR!!!)")
		} else {
			t.Log("Success")
		}
	}

	if l, err := w.Get(c, Test{Blah: "blah"}); err != nil {
		t.Error(err)
	} else {
		t.Log("Get Filtered Test 3")
		t.Log(PrintModel(l))
		if len(l) != 2 {
			t.Error("GetFiltered failed (ERROR!!!)")
		} else {
			t.Log("Success")
		}
	}

	if x, err := w.FindOne(c, test1); err != nil {
		t.Error(err)
	} else {
		t.Log("FindOneTest 1")
		t.Log("FindOne ")
		t.Log(PrintModel(x))

	}

	if x, err := w.FindOne(c, test2); err != nil {
		t.Error(err)
	} else {
		t.Log("FindOne Test 2")
		t.Log("FindOne ")
		t.Log(PrintModel(x))

	}

	if x, err := w.FindOneWithOr(c, Test{Name: "Test 3"}, Test{Name: "Test 1"}, Test{}); err != nil {
		t.Error(err)
	} else {
		t.Log("FindOneWithOrTest 1")
		t.Log("FindOneWithOr ")
		t.Log(PrintModel(x))
		if !(x.(*Test).Name == "Test 1") {
			t.Error("FindOneWithOr Failed")
		}

	}

	x, err := w.FindOneAndUpdate(c, test1, test3)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("FindOneAndUpdate Test 1")
		t.Log("FindOneAndUpdate ")
		t.Log(PrintModel(x))
		if !(x.(*Test).Message == "Test 1 Edited") {
			t.Error("FindOneAndUpdate Failed")
		}
		test3.Message = "Test 2 Edited"
	}

	y, err := w.FindOneAndUpdate(c, test2, test3)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("FindOneAndUpdate Test 2")
		t.Log("FindOneAndUpdate ")
		t.Log(PrintModel(y))
	}

	if n, err := w.DeleteOne(c, &Test{Name: "Test 1"}); err != nil {
		t.Error(err)
	} else {
		t.Log("Delete One Test 1")
		t.Log("Delete One Success : ")
		t.Log(n)
	}

	if n, err := w.DeleteOne(c, &Test{Name: "Test 2"}); err != nil {
		t.Error(err)
	} else {
		t.Log("Delete One Test 2")
		t.Log("Delete One Success : ")
		t.Log(n)
	}
}
