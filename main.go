package main

import (
	"log"
	"net/http"

	"srv/models/db"
	"srv/usables"
	"srv/views"
	M "srv/views/viewContext"
)

func main() {
	session, err := db.Init("proyecto")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	http.HandleFunc("/ingresar", M.MakeView(views.Login, "log_solicitud", "allow_crossSite"))

	log.Printf("Escuchando en %s...\n", usables.Dominio)
	err = http.ListenAndServe(":10443", nil) //TLS(":10443", "./server.crt", "./server.key", nil)
	log.Fatal(err)
}
