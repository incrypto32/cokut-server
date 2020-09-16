package models

import (
	"errors"
	"fmt"

	"github.com/incrypt0/cokut-server/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Phone   string             `json:"phone,omitempty" bson:"phone,omitempty" `
	LogoUrl string             `json:"logo,omitempty" bson:"logo,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Address string             `json:"address,omitempty" bson:"address,omitempty"`
	Type    string             `json:"type,omitempty" bson:"type,omitempty"`
}

func (r *Restaurant) GetModelData() string {
	return services.PrintModel(r)
}

func (r *Restaurant) Validate() error {
	fmt.Println(services.PrintModel(r))

	if r.Name == "" || (len(r.Phone) < 10) || r.Address == "" {
		return errors.New("Not Validated")
	}
	return nil
}
