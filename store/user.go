package store

import (
	"errors"
	"fmt"
	"log"

	"github.com/incrypt0/cokut-server/models"
)

// Function to insert users into userCollection
func (s *Store) InsertUser(u *models.User) (id string, err error) {
	var l interface{}

	fmt.Println("Test 1")
	//  Getting the user colection
	c := s.uc

	// Basic Validation
	if err = u.ValidateBasic(); err != nil {
		return id, err
	}

	// Check if email is null
	if u.Email != "" {

		if err = u.ValidateEmail(); err != nil {

			return id, err
		}
		fmt.Println("HERE")
		l, err = s.w.FindOneWithOr(c, models.User{Email: u.Email}, models.User{Phone: u.Phone})

	} else {

		l, err = s.w.FindOne(c, models.User{Phone: u.Phone})
	}

	fmt.Println(l)

	if err != nil {
		if err.Error() != "NIL" {
			log.Println(err)
			return id, err
		}

	}
	if l != nil {
		return id, errors.New("DETAILS_EXIST")
	}

	return s.w.Add(c, u)
}

// Check User existence
func (s *Store) CheckUserExistence(phone string) (bool, error) {
	var val bool = true
	fmt.Println("CheckUser called with phone : ", phone)
	c := s.uc
	filter := models.User{Phone: phone}
	l, err := s.w.FindOne(c, filter)

	if err != nil {

		if err.Error() == "NIL" {

			val = false
		} else {
			return false, err
		}
	}
	if l != nil {
		val = true
	}

	fmt.Println(phone, " exists : ", val)
	return val, nil

}
