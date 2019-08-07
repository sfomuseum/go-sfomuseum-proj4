# go-epsg

Go package for looking up EPSG codes.

## Install

You will need to have both `Go` (specifically version [1.12](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make tools
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Example

```
package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-epsg"
	"log"
)

func main() {

	flag.Parse()

	for _, str_code := range flag.Args(){
		def, _ := epsg.LookupString(str_code)
		fmt.Println(def)
	}
}
```

_Error handling removed for the sake or brevity._

## See also

* https://github.com/OSGeo/proj.4/
* https://raw.githubusercontent.com/OSGeo/proj.4/master/data/epsg