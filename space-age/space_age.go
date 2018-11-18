// space implements functions about time in space.
package space

type Planet string

var daySeconds = 24.0 * 60.0 * 60.0
var orbitalDayPeriods = map[Planet]float64{
	"Mercury": 87.97,
	"Venus":   224.70,
	"Earth":   365.26,
	"Mars":    686.98,
	"Jupiter": 4332.82,
	"Saturn":  10755.70,
	"Uranus":  30687.15,
	"Neptune": 60190.03,
}

// Age returns the number of planet-years that have elapsed, given a number of
// Earth-seconds and the name of the planet (case-sensitive). Returns Inf if
// planet doesn't exist in our solar system.
func Age(seconds float64, planet Planet) float64 {
	orbitalDays := orbitalDayPeriods[planet]
	return seconds / (orbitalDays * daySeconds)
}
