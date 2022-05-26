package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre,omitempty" json:"nombre"`
	Apellido        string             `bson:"apellido,omitempty" json:"apellido"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento,omitempty" json:"fechaNacimiento"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avagtar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biography       string             `bson:"biography" json:"biography,omitempty"`
	Location        string             `bson:"location" json:"location,omitempty"`
	WebSite         string             `bson:"website" json:"website,omitempty"`
}
