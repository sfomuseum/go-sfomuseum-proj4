package proj4

import (
	"errors"
	"io/ioutil"
	_ "log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type CLIProjector struct {
	Projector
}

func NewCLIProjector() (Projector, error) {
	p := CLIProjector{}
	return &p, nil
}

func (p *CLIProjector) Convert(c *Coordinate, src Projection, dest Projection) (*Coordinate, error) {

	// echo '6016820.327411 2048793.166111 30.56' | cs2cs -f "%.6f" +proj=lcc +lat_1=38.43333333333333 +lat_2=37.06666666666667 +lat_0=36.5 +lon_0=-120.5 +x_0=2000000.0001016 +y_0=500000.0001016001 +datum=NAD83 +units=us-ft +no_defs +to +proj=longlat +datum=WGS84 +no_defs

	fh, err := ioutil.TempFile("", "proj4")

	if err != nil {
		return nil, err
	}

	fname := fh.Name()

	defer func() {
		os.Remove(fname)
	}()

	args := []string{
		"echo",
		strconv.FormatFloat(c.X, 'f', 10, 32),
		strconv.FormatFloat(c.Y, 'f', 10, 32),
		strconv.FormatFloat(c.Z, 'f', 10, 32),
		"|",
		"/usr/local/bin/cs2cs",
		"-f",
		"\"%.6f\"",
		string(src),
		"+to",
		string(dest),
	}

	str_args := strings.Join(args, " ")
	// log.Println(str_args)

	fh.Write([]byte(str_args))
	fh.Close()

	cmd := exec.Command("sh", fname)
	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	str_out := strings.TrimSpace(string(out))
	parts := strings.Fields(str_out)

	if len(parts) != 3 {
		return nil, errors.New("Unknown output")
	}

	x, err := strconv.ParseFloat(parts[0], 32)

	if err != nil {
		return nil, err
	}

	y, err := strconv.ParseFloat(parts[1], 32)

	if err != nil {
		return nil, err
	}

	z, err := strconv.ParseFloat(parts[2], 32)

	if err != nil {
		return nil, err
	}

	c2, err := NewCoordinate(x, y, z)

	if err != nil {
		return nil, err
	}

	return c2, nil
}
