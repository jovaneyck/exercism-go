package space

type Planet = string

var secondsInOneEarthYear = 31557600.0
var orbitalPeriodsInEarthYears = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func round(n float64) float64 {
	return float64(int64(n/0.01+0.5)) * 0.01
}

func Age(seconds float64, planet string) float64 {
	secondsInAYearOnPlanet := secondsInOneEarthYear * orbitalPeriodsInEarthYears[planet]
	return round(seconds / secondsInAYearOnPlanet)
}
