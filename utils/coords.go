package utils

import (
	"math"
	"math/rand"
)

// GenerateRandomCoordinatesWithinRadius genera un lat y long tal que la ditancia sea ~5km del inicio
func GenerateRandomCoordinatesWithinRadius(startLat, startLon, radiusKm float64) (endLat, endLon float64) {
	// Aproximadamente 1 grado de latitud = 111 km
	radiusInDegrees := radiusKm / 111.0

	// Generamos un ángulo aleatorio
	angle := rand.Float64() * 2 * math.Pi
	// Generamos una distancia aleatoria dentro del radio
	distance := rand.Float64() * radiusInDegrees

	// Calculamos las nuevas coordenadas
	endLat = startLat + (distance * math.Cos(angle))
	endLon = startLon + (distance * math.Sin(angle) / math.Cos(startLat*math.Pi/180))

	return endLat, endLon
}
