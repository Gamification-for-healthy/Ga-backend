package model

import "time"

// untuk connect db
type DBinfo struct {
	DBString string
	DBName   string
}

// untuk password hash
type ResponseEncode struct {
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
}

type Payload struct {
	ID  string    `json:"id"`
	Exp time.Time `json:"exp"`
	Iat time.Time `json:"iat"`
	Nbf time.Time `json:"nbf"`
}
