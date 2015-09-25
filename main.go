package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/munnerz/haproxy/api"
	"github.com/munnerz/haproxy/api/writer"
)

var (
	filename = flag.String("file", "input.json", "Specify the JSON file to convert to HAProxy config")
)

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*filename)

	if err != nil {
		log.Fatalf("Error reading file: %s", err.Error())
	}

	input := api.JSONInput{}

	err = json.Unmarshal(b, &input)

	if err != nil {
		log.Fatalf("Error parsing json file: %s", err.Error())
	}

	output := []byte{}

	w := writer.NewWriter(&output)

	n, err := writer.WriteWriteable(w, &input)

	if err != nil {
		log.Fatalf("Error writing writeable to writer after %d bytes: %s", n, err.Error())
	}

	log.Printf("Config: \n%s", string(output))
}
