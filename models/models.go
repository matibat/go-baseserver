package models

import "gopkg.in/mgo.v2/bson"

//import "math/big"

// Session -> modelo
type Session struct {
	Sid       []byte //big.Int
	Tok       []byte //big.Int
	IDUsuario bson.ObjectId
}

// User -> modelo
type User struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	User string
	Pass []byte
	Salt []byte
	//Email string
	Stat byte
}
