package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-sfomuseum-proj4"
	"log"
	"strconv"
)

func main() {

	from := flag.String("from", "", "...")
	to := flag.String("to", "", "...")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("Insufficient arguments")
	}

	x, err := strconv.ParseFloat(args[0], 32)

	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.ParseFloat(args[1], 32)

	if err != nil {
		log.Fatal(err)
	}

	z := 0.0

	if len(args) > 2 {

		cz, err := strconv.ParseFloat(args[2], 32)

		if err != nil {
			log.Fatal(err)
		}

		z = cz
	}

	c, err := proj4.NewCoordinate(x, y, z)

	if err != nil {
		log.Fatal(err)
	}

	pr, err := proj4.NewProj4Projector()

	if err != nil {
		log.Fatal(err)
	}

	src, err := proj4.NewProjectionFromString(*from)

	if err != nil {
		log.Fatal(err)
	}

	target, err := proj4.NewProjectionFromString(*to)

	if err != nil {
		log.Fatal(err)
	}

	c2, err := pr.Convert(c, src, target)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c2)
}
