package proj4

// https://github.com/OSGeo/proj.4/blob/master/data/epsg
// https://raw.githubusercontent.com/OSGeo/proj.4/master/data/epsg

const (
	EPSG_4326 Projection = "+proj=longlat +ellps=WGS84 +datum=WGS84 +no_defs"
	EPSG_2227 Projection = "+proj=lcc +lat_1=38.43333333333333 +lat_2=37.06666666666667 +lat_0=36.5 +lon_0=-120.5 +x_0=2000000.0001016 +y_0=500000.0001016001 +datum=NAD83 +units=us-ft +no_defs"
)
