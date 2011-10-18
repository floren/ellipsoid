package ellipsoid

import "fmt"

func main() {
	lat1, lon1 := 37.619002, -122.374843
	lon2, lat2 := 33.942536, -118.408074
	{
		geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("1 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("2 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Radians, ellipsoid.Meter, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_not_symmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("3 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Radians, ellipsoid.Meter, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("4 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Kilometer, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("5 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Foot, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("6 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Foot, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		lon2, lat2 := geo1.At(lat1, lon1, 2000.0, 45.0)
		fmt.Printf("7 lon  = %v lat = %v\n", lon2, lat2)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		lon2, lat2 := geo1.At(lat1, lon1, 2000.0, 45.0)
		fmt.Printf("8 lon  = %v lat = %v\n", lon2, lat2)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
		dist, bear := geo1.At(90.0, 90.0, 1000.0, 90.0)
		fmt.Printf("9 dist  = %v bear = %v\n", dist, bear)
	}
}