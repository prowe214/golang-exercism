// Package space creates functions for determining age on different planets
package space

// Planet is struct for planets
type Planet string

// earthYearSecs stores earth year length in seconds
var earthYearSecs float64 = 31557600

// yearLength stores each planet's year length relative to earth
var yearLength = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age Determines age on different planets
func Age(ageInSeconds float64, planet Planet) float64 {

	// Calculates this planet's year length in seconds
	planetYearInSeconds := yearLength[planet] * earthYearSecs

	// Calculates your age on the new planet in years
	planetAge := ageInSeconds / planetYearInSeconds

	return planetAge
}
