package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-epsg"
	"log"
	"net/http"
	"time"
)

func main() {

	epsg_data := flag.String("data", "https://raw.githubusercontent.com/OSGeo/proj.4/master/data/epsg", "The URL of your EPSG source data")

	flag.Parse()

	rsp, err := http.Get(*epsg_data)

	if err != nil {
		log.Fatal(err)
	}

	defer rsp.Body.Close()

	defs, err := epsg.MakeDefinitions(rsp.Body)

	if err != nil {
		log.Fatal(err)
	}

	enc_defs, err := json.Marshal(defs)

	if err != nil {
		log.Fatal(err)
	}

	ts := time.Now()

	fmt.Printf("%s\n\n", "package epsg")

	fmt.Printf("/* This file was generated by robots (%s) at %s */\n\n", "cmd/mk-definitions.go", ts.UTC())
	fmt.Printf("const Definitions string = `%s`", string(enc_defs))
}
