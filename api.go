package proj4

/*
#cgo LDFLAGS: -lproj
#include <string.h>
#include <proj_api.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Proj4Error struct {
	retval int
}

func (e *Proj4Error) Error() string {
	return fmt.Sprintf("Proj4 Error: %d", e.retval)
}

type Proj4Projector struct {
	Projector
}

func NewProj4Projector() (Projector, error) {
	p := Proj4Projector{}
	return &p, nil
}

func (p *Proj4Projector) Convert(c *Coordinate, src Projection, dest Projection) (*Coordinate, error) {

	if src == EPSG_4326 {
		c = c.ToRadians()
	}

	// With cs2cs being the command line interface to the old API, and cct being the same for the new, this example
	// of doing the same thing in both world views will should give an idea of the differences:
	// https://github.com/OSGeo/proj.4/blob/811fc90adeb6782dded4ef5c2ab58131e8399c67/docs/source/development/migration.rst
	// https://github.com/OSGeo/proj.4/blob/811fc90adeb6782dded4ef5c2ab58131e8399c67/src/cs2cs.c

	fromProj := C.CString(string(src))
	toProj := C.CString(string(dest))

	defer func() {
		C.free(unsafe.Pointer(fromProj))
		C.free(unsafe.Pointer(toProj))
	}()

	fpj := C.pj_init_plus(fromProj)
	tpj := C.pj_init_plus(toProj)

	lng := C.double(c.X)
	lat := C.double(c.Y)
	ele := C.double(c.Z)

	retval := int(C.pj_transform(fpj, tpj, 1, 1, &lng, &lat, &ele))

	C.pj_free(fpj)
	C.pj_free(tpj)

	if retval != 0 {
		return nil, &Proj4Error{retval}
	}

	rsp := &Coordinate{
		X: float64(lng),
		Y: float64(lat),
		Z: float64(ele),
	}

	if dest == EPSG_4326 {
		rsp = rsp.ToDegrees()
	}

	return rsp, nil
}
