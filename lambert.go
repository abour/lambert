package main

import (
	"math"
)

func Lambert93ToWGPS(lambertE float64, lambertN float64) (float64, float64) {
	GRS80E := 0.081819191042816
	LONG_0 := 3.
	XS := 700000.
	YS := 12655612.0499
	n := 0.7256077650532670
	C := 11754255.4261

	delX := lambertE - XS
	delY := lambertN - YS
	gamma := math.Atan(-delX / delY)
	R := math.Sqrt(delX*delX + delY*delY)
	latiso := math.Log(C/R) / n

	last := math.Tanh(latiso + GRS80E*math.Atanh(GRS80E*math.Sin(1)))
	precision := 6
	for i := 0; i < precision; i++ {
		new := math.Tanh(latiso + GRS80E*math.Atanh(GRS80E*last))
		last = new
	}

	longRad := math.Asin(last)
	latRad := gamma/n + LONG_0/180*math.Pi

	var longitude = latRad / math.Pi * 180
	var latitude = longRad / math.Pi * 180

	return longitude, latitude
}
