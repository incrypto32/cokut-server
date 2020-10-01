package store

import (
	"errors"
	"log"

	"github.com/incrypt0/cokut-server/models"
)

//InsertUser Function to insert users into userCollection
func (s *Store) InsertUser(u *models.User) (id string, err error) {
	var l interface{}

	//  Getting the user colection
	c := s.uc

	// Basic Validation
	if err = u.ValidateBasic(); err != nil {
		return id, err
	}

	l, err = s.w.FindOneWithOr(c, models.User{Email: u.Email},
		models.User{Phone: u.Phone},
		models.User{GID: u.GID},
		models.User{UID: u.UID})

	if err != nil {
		if err.Error() != "NIL" {
			log.Println(err)
			return id, errors.New("ERROR")
		}
	}

	if l != nil {
		return id, errors.New("DETAILS_EXIST")
	}

	return s.w.Add(c, u)
}

//InsertUser Function to insert users into userCollection
func (s *Store) AddUserAddress(uid string, address models.Address) (user *models.User, err error) {
	var i interface{}

	//  Getting the user colection
	c := s.uc

	if err != nil {
		return user, errors.New("ERROR")
	}

	i, err = s.w.FindOneAndPush(c, models.User{UID: uid}, address, "address")

	if err != nil {
		if err.Error() != "NIL" {
			log.Println(err)
			return user, errors.New("ERROR")
		}
	}

	user = i.(*models.User)

	return user, err
}

//InsertUser Function to insert users into userCollection
func (s *Store) RemoveUserAddress(uid string, address models.Address) (user *models.User, err error) {
	var i interface{}

	//  Getting the user colection
	c := s.uc

	if err != nil {
		return user, errors.New("ERROR")
	}

	i, err = s.w.FindOneAndPull(c, models.User{UID: uid}, address, "address")

	if err != nil {
		if err.Error() != "NIL" {
			log.Println(err)
			return user, errors.New("ERROR")
		}
	}

	user = i.(*models.User)

	return user, err
}

// GetUser .
func (s *Store) GetUser(uid string) (l interface{}, err error) {
	return s.w.FindOne(s.uc, models.User{UID: uid})
}

//CheckUserPhoneExistence  checks whether the user exists with a phone
func (s *Store) CheckUserPhoneExistence(phone string) (bool, error) {
	var val bool = true

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

	return val, nil
}

//CheckUserPhoneExistenceByGID  checks whether the user exists with a phone
func (s *Store) CheckUserExistenceByGID(gid string) (bool, error) {
	var val bool = true

	c := s.uc
	filter := models.User{GID: gid}
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

	return val, nil
}

//CheckUserExistence checks whether the user exists based on email and phone
func (s *Store) CheckUserExistence(phone string, email string) (bool, error) {
	var val bool

	c := s.uc

	var l interface{}

	var err error

	// Check if email is null
	if email != "" {
		l, err = s.w.FindOneWithOr(c, models.User{Email: email}, models.User{Phone: email})
	} else {
		l, err = s.w.FindOne(c, models.User{Phone: phone})
	}

	if err != nil {
		if err.Error() == "NIL" {
			val = false
		} else {
			log.Println(err)
		}
	}

	if l != nil {
		val = true
	}

	return val, err
}

//CheckUserExistenceByUID checks whether the user exists based on UID
func (s *Store) CheckUserExistenceByUID(uid string) (bool, error) {
	var val bool

	var l interface{}

	var err error

	c := s.uc

	l, err = s.w.FindOne(c, models.User{UID: uid})

	if err != nil {
		if err.Error() == "NIL" {
			err = nil
			val = false
		} else {
			log.Println(err)
		}
	}

	if l != nil {
		val = true
	}

	return val, err
}
