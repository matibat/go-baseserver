package views

import (
	"log"
	"net/http"
	"servidor/src/auth"
	"servidor/src/forms"
)

// Login -> Iniciar sesión
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials forms.Credentials
	var status forms.Status
	err := readJSON(&credentials, r)
	if err != nil {
		log.Printf("[Login] Error interno: %s\n", err)
		status.Code = 2 // Error interno
	} else {
		err = auth.Authenticate(credentials.User, credentials.Pass)
		if err != nil {
			log.Print(err)
			status.Code = 1 // Usuario y contraseña no coinciden
		} else {
			_, err = auth.Login(credentials.User)
			if err != nil {
				log.Printf("[Login] Error interno: %s\n", err)
				status.Code = 2 // Error interno
			}
		}
	}
	err = writeJSON(status, w)
	if err != nil {
		log.Fatalf("[Login] Error interno fatal: %s\n", err) // A la mierda todo. Cagamos. :v
	}
}

