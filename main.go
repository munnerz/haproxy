package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"

	"github.com/munnerz/haproxy/api"
	"github.com/munnerz/haproxy/api/writer"

	"github.com/gorilla/mux"
)

const (
	port = 10000
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", CreateHandler).Methods("POST")

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		log.Errorf("Error reading request body: %s", err.Error())
		return
	}

	input := &api.JSONInput{}

	err = json.Unmarshal(b, input)

	if err != nil {
		log.Errorf("Error decoding JSON input: %s", err.Error())
		return
	}

	_, err = writer.WriteWriteable(w, input)

	if err != nil {
		log.Errorf("Error writing response: %s", err.Error())
		return
	}

}
