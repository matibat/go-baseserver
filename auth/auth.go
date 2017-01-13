package auth

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	//"math/big"
	"srv/models"
	"srv/models/db"
	"srv/usables"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	//mgo "gopkg.in/mgo.v2"
)

//var maxSession = usables.GetBigint("________________________________________________________________________________________________________________________________")

// Authenticate -> Verifica si las credenciales del usuario son válidas
func Authenticate(user string, pass []byte) error {
	log.Printf("Authenticating '%s'...\n", user)
	var userdata models.User
	q := db.Session.Find(bson.M{"user": user})
	err := q.One(&userdata)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdata.Pass), append(pass, userdata.Salt...))
	if err != nil {
		return err
	}
	log.Printf("'%s' sucessfully authenticated.", user)
	return nil
}

// CreateUser -> Crea un nuevo usuario
func CreateUser(username string, password []byte) (models.User, error) {
	var newuser models.User
	var salt = usables.GetRandBytes(128)
	pass, err := bcrypt.GenerateFromPassword(append(password, salt...), 1)
	if err != nil {
		return models.User{}, err
	}
	newuser = models.User{
		User: username,
		Pass: pass,
		Salt: salt,
		Stat: 1,
	}
	err = db.User.Insert(newuser)
	if err != nil {
		return models.User{}, err
	}
	return newuser, nil
}

// Login -> Loguea al usuario, inicia una sesión
func Login(user string) (models.Session, error) {
	log.Printf("Logging in '%s'...", user)
	results := db.User.Find(bson.M{"user": user})
	nresults, err := results.Count()
	if err != nil {
		return models.Session{}, err
	}
	if nresults == 1 {
		var userdata models.User
		var sessiondata models.Session
		results.One(&userdata)
		err = errors.New("")
		for err != nil {
			sessiondata = models.Session{
				Sid:       usables.GetRandBytes(128),
				Tok:       usables.GetRandBytes(128),
				IDUsuario: userdata.ID,
			}
			err = db.Session.Insert(sessiondata)
		}
		return sessiondata, nil
	}
	return models.Session{}, errors.New("components (auth.go): User not found")
}

// Logout -> Cierra la sesión
func Logout(sid *big.Int) error {
	sidbin := sid.Bytes()
	q := db.Session.Find(bson.M{"sid": sidbin})
	nresults, err := q.Count()
	if err != nil {
		return err
	}
	fmt.Printf("Resultados: %d\n", nresults)
	if nresults == 1 {
		err = db.Session.Remove(bson.M{"sid": sidbin})
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("components (auth.go): Usuario no logueado _")
}
