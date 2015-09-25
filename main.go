package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/munnerz/haproxy/api"
	"github.com/munnerz/haproxy/api/writer"
)

var (
	filename = flag.String("file", "input.json", "Specify the JSON file to convert to HAProxy config")
)

func writeString(w io.Writer, s string) (n int, err error) {
	return w.Write([]byte(fmt.Sprintf("%s\n", s)))
}

func writeWriteable(w io.Writer, o api.Writeable) {
	for _, s := range o.Strings() {
		writeString(w, s)
	}
}

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

	writeWriteable(w, &input)

	log.Printf("Config: \n%s", string(output))
}
