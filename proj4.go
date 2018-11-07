package proj4

import (
	"errors"
	"fmt"
	"strings"
)

const (
	radians float64 = 0.017453292519943295
	degrees float64 = 57.29577951308232
)

type Projection string

func NewProjectionFromString(name string) (Projection, error) {

	name = strings.Replace(name, "epsg:", "", -1)

	var p Projection
	var err error

	switch name {
	case "4326":
		p = EPSG_4326
	case "2227":
		p = EPSG_2227
	default:
		err = errors.New("Invalid or unsupported EPSG string")
	}

	return p, err
}

type Projector interface {
	Convert(*Coordinate, Projection, Projection) (*Coordinate, error)
}

type Coordinate struct {
	X float64
	Y float64
	Z float64
}

func NewCoordinate(args ...float64) (*Coordinate, error) {

	var c *Coordinate
	var err error

	switch len(args) {

	case 3:
		c = &Coordinate{
			X: args[0],
			Y: args[1],
			Z: args[2],
		}
	case 2:
		c = &Coordinate{
			X: args[0],
			Y: args[1],
			Z: 0.0,
		}
	default:
		err = errors.New("Invalid X,Y,Z arguments")
	}

	return c, err
}

func (c *Coordinate) ToDegrees() *Coordinate {

	return &Coordinate{
		X: c.X * degrees,
		Y: c.Y * degrees,
		Z: c.Z,
	}
}

func (c *Coordinate) ToRadians() *Coordinate {

	return &Coordinate{
		X: c.X * radians,
		Y: c.Y * radians,
		Z: c.Z,
	}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%.6f %.6f %.6f", c.X, c.Y, c.Z)
}
