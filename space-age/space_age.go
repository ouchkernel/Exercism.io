package space

type Planet string

const earthSeconds = 31557600

func Age(seconds float64, p Planet) float64 {
	var earthYear float64
	switch p {
	case "Earth":
		earthYear = 1
	case "Mercury":
		earthYear = 0.2408467
	case "Venus":
		earthYear = 0.61519726
	case "Mars":
		earthYear = 1.8808158
	case "Jupiter":
		earthYear = 11.862615
	case "Saturn":
		earthYear = 29.447498
	case "Uranus":
		earthYear = 84.016846
	case "Neptune":
		earthYear = 164.79132
	}
	return seconds / (earthSeconds * earthYear)
}
