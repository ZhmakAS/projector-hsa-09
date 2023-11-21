package main

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Id        int64     `json:"id" bson:"id" db:"id" fake:"{number:1,10}"`
	FirstName string    `json:"first_name" db:"first_name" fake:"{firstname}"`
	LastName  string    `json:"last_name" db:"last_name" fake:"{lastname}"`
	Phone     string    `json:"phone" db:"phone" fake:"{phone}"`
	BirthDate time.Time `json:"birth_date" db:"birth_date"`
}

func NewUser() User {
	var newUser User
	gofakeit.Struct(&newUser)
	newUser.BirthDate = gofakeit.DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))
	return newUser
}
