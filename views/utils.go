package views

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// readJSON -> Lee datos recibidos
func readJSON(aStruct interface{}, r *http.Request) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, aStruct)
	if err != nil {
		return err
	}
	return nil
}

func writeJSON(aStruct interface{}, w http.ResponseWriter) error {
	data, err := json.Marshal(aStruct)
	if err != nil {
		return err
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
	return nil
}
