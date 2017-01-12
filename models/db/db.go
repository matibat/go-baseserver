package db

import (
	mgo "gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

// DB -> Puntero de referencia al controlador mongodb
var DB *mgo.Database

// Session -> Conlección
var Session *mgo.Collection

// User -> Colección
var User *mgo.Collection

// Init -> Inicializa la conexion con la base
func Init(base string) (*mgo.Session, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	DB = session.DB(base)
	Session = DB.C("session")
	User = DB.C("user")

	Session.EnsureIndex(mgo.Index{
		Key:    []string{"id"},
		Unique: true,
	})
	User.EnsureIndex(mgo.Index{
		Key:    []string{"user"},
		Unique: true,
	})
	return session, nil
}
