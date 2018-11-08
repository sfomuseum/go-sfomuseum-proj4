package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-epsg"
	"log"
)

func main() {

	flag.Parse()

	for _, str_code := range flag.Args() {

		def, ok := epsg.LookupString(str_code)

		if !ok {
			log.Fatal("Invalid code")
		}

		fmt.Println(def)
	}
}
