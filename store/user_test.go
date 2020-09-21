package store

import (
	"testing"

	"github.com/incrypt0/cokut-server/models"

	"github.com/incrypt0/cokut-server/workers"
)

// func TestUser(t *testing.T) {
// 	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())

// 	a := make([]string, 3)
// 	a[0] = "Vazahppully House"
// 	user := models.User{
// 		Name:    "Krish",
// 		Address: a,
// 		Email:   "vpkrishnanand@gmail.com",
// 		Admin:   true,
// 		Phone:   "7034320441",
// 		UID:     "1",
// 	}

// 	if id, err := s.InsertUser(&user); err != nil {
// 		if err.Error() == "DETAILS_EXIST" {
// 			t.Log("Details Exist")
// 		} else {
// 			t.Error("ERROR >>>>>>>>>>> : ", err)
// 		}
// 	} else {
// 		t.Log("Insert User Success : ", id)
// 	}

// 	if val, err := s.CheckUserExistence("7034320441", "vpkrishnanand@gmail.com"); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Phone Exists (Test 1) : ", val)
// 	}

// 	if val, err := s.CheckUserExistence("9847859164", "blah@gmail.com"); err != nil {
// 		if err.Error() != "NIL" {
// 			t.Error("ERROR >>>>>>>>>>> : ", err)
// 		}
// 	} else {
// 		t.Log("User Exists (Test 2) : ", val)
// 	}
// 	if val, err := s.CheckUserExistence("7034320441", "blah@gmail.com"); err != nil {
// 		if err.Error() != "NIL" {
// 			t.Error("ERROR >>>>>>>>>>> : ", err)
// 		}
// 	} else {
// 		t.Log("User Exists (Test 1) : ", val)
// 	}

// 	if val, err := s.CheckUserExistence("9847859164", "vpkrishnanand@gmail.com"); err != nil {
// 		if err.Error() != "NIL" {
// 			t.Error("ERROR >>>>>>>>>>> : ", err)
// 		}
// 	} else {
// 		t.Log("User Exists (Test 2) : ", val)
// 	}

// 	if val, err := s.CheckUserPhoneExistence("7034320441"); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Phone Exists (Test 1) : ", val)
// 	}

// 	if val, err := s.CheckUserPhoneExistence("9847859164"); err != nil {
// 		if err.Error() != "NIL" {
// 			t.Error("ERROR >>>>>>>>>>> : ", err)
// 		}
// 	} else {
// 		t.Log("Phone Exists (Test 2) : ", val)
// 	}
// }

// func TestRestaurant(t *testing.T) {
// 	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())

// 	id1, _ := primitive.ObjectIDFromHex("5f689aa84770a61206b9095b")
// 	id2, _ := primitive.ObjectIDFromHex("5ec0115c23ca01bb29ea922f")

// 	res1 := models.Restaurant{
// 		ID:      id1,
// 		Name:    "Ambadi Hotel",
// 		Address: "Sasikutan Lane",
// 		Email:   "hai@gmail.com",
// 		LogoUrl: "imgur.com/blah",
// 		Phone:   "9847859163",
// 		Type:    "home",
// 	}

// 	res2 := models.Restaurant{
// 		ID:      id2,
// 		Name:    "Ambadi Hotel 2",
// 		Address: "Sasikutan Lane 2",
// 		Email:   "hai2@gmail.com",
// 		LogoUrl: "imgur.com/blah2",
// 		Phone:   "9847859162",
// 		Type:    "reg",
// 	}

// 	t.Log("<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>")
// 	t.Log(id1, id2)

// 	if id, err := s.InsertRestaurant(&res1); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Insert Restaurant Test 1")
// 		t.Log("Insert Restaurant Success : ", id)
// 	}

// 	if id, err := s.InsertRestaurant(&res2); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Insert Restaurant Test 2")
// 		t.Log("Insert Restaurant Success : ", id)
// 	}

// 	if l, err := s.GetAllRestaurants(); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Get All Restaurants Test 1")
// 		t.Log(utils.ModelToString(l))
// 	}

// 	if l, err := s.GetAllHomeMade(); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Get All HomeMade Restaurants Test 1")
// 		t.Log(utils.ModelToString(l))
// 	}
// }

// func TestMeals(t *testing.T) {
// 	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())
// 	id1, _ := primitive.ObjectIDFromHex("5f689aa84770a61206b9095b")
// 	id2, _ := primitive.ObjectIDFromHex("5ec0115c23ca01bb29ea922f")

// 	m1 := models.Meal{

// 		Name:         "Cheesy Pizza",
// 		DisplayPrice: 85.2,
// 		Price:        75.0,
// 		IsVeg:        false,
// 		RID:          id1.Hex(),
// 		Spicey:       true,
// 	}

// 	m2 := models.Meal{

// 		Name:         "Chicken Kabab",
// 		DisplayPrice: 65.2,
// 		Price:        77.0,
// 		IsVeg:        false,
// 		RID:          id2.Hex(),
// 		Spicey:       true,
// 	}

// 	var mid string

// 	if id, err := s.InsertMeal(&m1); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Insert Meal Test 1")
// 		t.Log("Insert Meal Success : ", id)
// 	}

// 	mid, err := s.InsertMeal(&m2)
// 	if err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Insert Meal Test 2")
// 		t.Log("Insert Meal Success : ", mid)
// 	}

// 	if id, err := s.InsertSpecial(mid); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Insert Special Test 1")
// 		t.Log("Insert Special Success : ", id)
// 	}

// 	if l, err := s.GetMealsByRestaurant(id1.Hex()); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("GetMeals By Restaurant Test 1")
// 		t.Log("GetMeals By Restaurant : ")
// 		t.Log(utils.ModelToString(l))
// 	}

// 	if l, err := s.GetSpecialMeals(); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Get Special Meals Test 1")
// 		t.Log("Get Special Meals : ")
// 		t.Log(utils.ModelToString(l))
// 	}

// 	if l, err := s.GetSpiceyMeals(); err != nil {
// 		t.Error("ERROR >>>>>>>>>>> : ", err)
// 	} else {
// 		t.Log("Get Spicey Meals Test 1")
// 		t.Log("Get Spicey Meals Restaurant : ")
// 		t.Log(utils.ModelToString(l))
// 	}
// }

func TestUser1(t *testing.T) {
	var s *Store = NewStore("mctest", "uctest", "octest", "rctest", workers.New())
	_ = s.w.DropTest()
	a := make([]string, 3)
	a[0] = "Vazahppully House"
	user := &models.User{
		Name:    "Krish",
		Address: a,
		Email:   "vpkrishnanand@gmail.com",
		Admin:   true,
		Phone:   "7034320441",
		UID:     "1",
		GID:     "g1",
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
