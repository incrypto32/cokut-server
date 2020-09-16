package services

// import (
// 	"encoding/json"
// 	"testing"
// 	"time"
// )

// func Test(t *testing.T) {
// 	t.Log(time.Now().Unix())
// }
func blah() {}

// // DbServicesTest

// func TestDbSerivces(t *testing.T) {
// 	type test struct {
// 		Name    string `bson:"name" bson:"name"`
// 		Phone   string `bson:"phone" bson:"phone"`
// 		Address string `bson:"address" bson:"address"`
// 		Email   string `bson:"email" bson:"email"`
// 		Typ     string `bson:"type" bson:"type"`
// 	}
// 	tes := &test{
// 		Name:    "Hari",
// 		Phone:   "7034320441",
// 		Address: "Blha Spot",
// 		Email:   "krish@gmail.com",
// 		Typ:     "reg",
// 	}

// 	t.Log("_______DbServicesTest_______")
// 	ConnectMongo()

// 	t.Log(" |       Add()       | ")
// 	if id, err := Add(C.TestCollection, tes); err != nil {
// 		t.Error("Error Occured In Add()")
// 		t.Error(err)

// 	} else {
// 		t.Log("Add() completed successfully with inserted id : ", id)
// 	}

// 	t.Log(" --------Testing GetAll()------- ")

// 	if l, err := GetAll(C.TestCollection, test{}); err != nil {
// 		t.Error("Error Occured In Add()")
// 		t.Error(err)

// 	} else {
// 		if l[0].(*test).Email == "" {
// 			t.Error("Get All Failed")
// 		}
// 		t.Log("Tested Name : ", l[0].(*test).Name)
// 	}
// 	t.Log("_____DbServicesTestEnded_____")

// }

// func TestFire(t *testing.T) {

// 	InitFire()
// 	client, err := app.Auth(ctx)

// 	if err != nil {
// 		t.Error("Firebase Error")
// 		t.Error(err)
// 		return
// 	}

// 	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjQ5YWQ5YmM1ZThlNDQ3OTNhMjEwOWI1NmUzNjFhMjNiNDE4ODA4NzUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20va29zcy00ODhhYyIsImF1ZCI6Imtvc3MtNDg4YWMiLCJhdXRoX3RpbWUiOjE1OTk4MDE1MjQsInVzZXJfaWQiOiIzdkFRQkdLVFUyUU5oc2hhcTk0VWRHeHprUm8yIiwic3ViIjoiM3ZBUUJHS1RVMlFOaHNoYXE5NFVkR3h6a1JvMiIsImlhdCI6MTU5OTgwMTUyNCwiZXhwIjoxNTk5ODA1MTI0LCJwaG9uZV9udW1iZXIiOiIrOTE3MDM0MzIwNDQwIiwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJwaG9uZSI6WyIrOTE3MDM0MzIwNDQwIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGhvbmUifX0.zHZ13FJeCYFF315CIcRen_JTf_a4yGgz-0Q-y_Ii9SkmvqYmdISctZdEv42BM2nPGNL9LGPZ0yRKCXe5lZA6b0sFgCaVHrIKhDClqYCze-lmMYX7z9LiVvjJHhi1Li0fYkgn3EeL5uzn5EeA_R2Vi3xR1EbgwWN8iZ08feWOrZ3mjOrfsHGtVPPBpbFUyZQkI82D8Rf59L99L3e9Sp6ZJHXCRCsdJJ_lQkEL57XNKzsB8a9Ujjw1BYqHQGZ5iyubqfdOe868DT5HSNFpOwUB9TCvq2f2bEisspIXpFW5gjljU1z2q76t1EroS-fyJqQ_YABEzbfCWQCvqAlhy7ddKw"

// 	token, err := client.VerifyIDToken(ctx, idToken)

// 	if err != nil {
// 		t.Error("Firebase Error")
// 		t.Error(err)
// 		return
// 	}
// 	b, _ := json.MarshalIndent(token, "", "  ")
// 	t.Log(string(b))
// }
