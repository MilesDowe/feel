package util

import (
	"math"
)

// Mean : calculates mean
func Mean(data []float64) float64 {
	var sum float64

	for i := 0; i < len(data); i++ {
		sum = sum + data[i]
	}
	return (sum / float64(len(data)))
}

// StdDev : calculates standard deviation
func StdDev(data []float64) float64 {
	var dataMean, variance, temp float64

	squaredData := make([]float64, len(data))

	// Get the mean
	dataMean = Mean(data)

	// Get distance from mean, then square each value
	for i := 0; i < len(data); i++ {
		temp = data[i] - dataMean
		squaredData[i] = temp * temp
	}

	// Get the variance
	variance = Mean(squaredData)
	return float64(math.Sqrt(float64(variance)))
}
