package viewContext

import (
	//"database/sql"
	"log"
	"net/http"
)

const methodOPTIONS = "OPTIONS"

type vContext struct {
	funcion func(http.ResponseWriter, *http.Request)
	listaMW []string
	// db      *mgo.Database
}

var listaMiddleware = map[string]func(http.ResponseWriter, *http.Request) int{
	"log_solicitud": func(w http.ResponseWriter, r *http.Request) int {
		if r.Method != methodOPTIONS {
			log.Printf("%s %s ", r.Method, r.Header.Get("Content-Length"))
		}
		return 0
	},
	"allow_crossSite": func(w http.ResponseWriter, r *http.Request) int {
		// log.Printf("allow_crossSite: Metodo: %s\n", r.Method)
		if r.Method == methodOPTIONS {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("ORIGIN"))
			w.Header().Set("Access-Control-Allow-Headers", "accept, content-type")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Write(nil)
			return 1
		}
		return 0
	},
}

func (M vContext) NuevoHandler(w http.ResponseWriter, r *http.Request) {
	var Estado int
	for _, funcionMiddleware := range M.listaMW {
		Estado = listaMiddleware[funcionMiddleware](w, r)
		if Estado != 0 {
			break
		}
	}
	if Estado == 0 {
		M.funcion(w, r)
	} else if r.Method != methodOPTIONS {
		log.Printf("|-> Un middleware respondió antes que el Handler.\n")
	}
}

// MakeView -> Devuelve un http.Handler aplicándole el middleware especificado
func MakeView(h func(http.ResponseWriter, *http.Request), m ...string) func(http.ResponseWriter, *http.Request) {
	var descripcion = vContext{
		funcion: h,
		listaMW: m,
		//db:      db,
	}
	return descripcion.NuevoHandler
}
