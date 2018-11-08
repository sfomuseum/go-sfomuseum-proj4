package epsg

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type EPSGTable map[int]string

var defs *sync.Map

var re_def *regexp.Regexp
var re_code *regexp.Regexp

func init() {

	re_def = regexp.MustCompile(`<(\d+)>\s*([^<]+)\s*<>`)

	// I am apparently too dumb to make this work... (20181108/thisisaaronland)
	// re_code = regexp.MustCompile(`(?:(epsg|EPSG)[\:_])?(\d+)`)

	var tbl *EPSGTable
	err := json.Unmarshal([]byte(Definitions), &tbl)

	if err != nil {
		panic(err)
	}

	defs = new(sync.Map)

	for code, def := range *tbl {
		go defs.Store(code, def)
	}
}

func Lookup(code int) (string, bool) {

	def, ok := defs.Load(code)

	if !ok {
		return "", false
	}

	return def.(string), ok
}

func LookupString(str_code string) (string, bool) {

	code, err := StringToCode(str_code)

	if err != nil {
		return "", false
	}

	return Lookup(code)
}

func StringToCode(str_code string) (int, error) {

	str_code = strings.ToUpper(str_code)

	if strings.HasPrefix(str_code, "EPSG") {
		str_code = strings.Replace(str_code, "EPSG", "", -1)
		str_code = strings.Replace(str_code, ":", "", -1)
		str_code = strings.Replace(str_code, "_", "", -1)
	}

	return strconv.Atoi(str_code)

	// see above

	/*
			if !re_code.MatchString(str_code) {
				return -1, errors.New("Invalid code")
			}

		m := re_code.FindStringSubmatch(str_code)
		return strconv.Atoi(m[0])
	*/
}

func MakeDefinitions(r io.Reader) (*map[int]string, error) {

	scanner := bufio.NewScanner(r)

	lookup := new(sync.Map)
	wg := new(sync.WaitGroup)

	for scanner.Scan() {

		wg.Add(1)

		ln := scanner.Text()

		go func(ln string, lookup *sync.Map, wg *sync.WaitGroup) {

			defer wg.Done()

			if strings.HasPrefix(ln, "#") {
				return
			}

			if !re_def.MatchString(ln) {
				log.Println("%s does not match", ln)
				return
			}

			m := re_def.FindStringSubmatch(ln)

			str_code := m[1]

			code, err := strconv.Atoi(str_code)

			if err != nil {
				log.Fatal("failed to parse code", str_code)
				return
			}

			def := strings.Trim(m[2], " ")
			lookup.Store(code, def)

		}(ln, lookup, wg)
	}

	err := scanner.Err()

	if err != nil {
		return nil, err
	}

	wg.Wait()

	spec := make(map[int]string)

	lookup.Range(func(key interface{}, value interface{}) bool {

		code := key.(int)
		def := value.(string)

		spec[code] = def
		return true
	})

	return &spec, nil
}
