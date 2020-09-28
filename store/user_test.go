package store

import (
	"testing"

	"github.com/incrypt0/cokut-server/models"
	"github.com/incrypt0/cokut-server/workers"
)

func TestUser1(t *testing.T) {
	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())
	_ = s.w.DropTest()
	a := make([]string, 3)
	a[0] = "Vazahppully House"
	user := &models.User{
		Name:  "Krish",
		Email: "vpkrishnanand@gmail.com",
		Admin: true,
		Phone: "7034320441",
		UID:   "1",
		GID:   "g1",
	}

	if id, err := s.InsertUser(user); err != nil {
		t.Log(err)
	} else {
		t.Log("user inserted : ", id)
	}

	if l, err := s.GetUser("1"); err != nil {
		t.Log(err)
	} else if l.(*models.User).UID != "1" {
		t.Error("ERROR")
	}
}

func TestUser2(t *testing.T) {
	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())

	if val, err := s.CheckUserExistence("7034320441", "vpkrishnanand@gmail.com"); err != nil {
		t.Error(err)
	} else if !val {
		t.Error("ERROR Expected true")
	}

	if val, err := s.CheckUserExistence("1234567890", "vpkrishnanand@gmail.com"); err != nil {
		t.Error(err)
	} else if !val {
		t.Error("ERROR Expected true")
	}

	if val, err := s.CheckUserExistence("7034320441", "a@gmail.com"); err != nil {
		if err.Error() != "NIL" {
			t.Error(err)
		}
	} else if !val {
		t.Error("ERROR Expected true")
	}

	if val, err := s.CheckUserExistence("1234567890", "a@gmail.com"); err != nil {
		if err.Error() != "NIL" {
			t.Error(err)
		}
	} else if val {
		t.Error("ERROR Expected true")
	}
}

func TestUser3(t *testing.T) {
	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())

	if val, err := s.CheckUserExistenceByUID("1"); err != nil {
		t.Error(err)
	} else if !val {
		t.Error("ERROR Expected true")
	}

	if val, err := s.CheckUserExistenceByGID("g1"); err != nil {
		t.Error(err)
	} else if !val {
		t.Error("ERROR Expected true")
	}

	if val, err := s.CheckUserPhoneExistence("7034320441"); err != nil {
		t.Error(err)
	} else if !val {
		t.Error("ERROR Expected true")
	}
}
func TestDrop(t *testing.T) {
	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())

	_ = s.w.DropTest()
}
