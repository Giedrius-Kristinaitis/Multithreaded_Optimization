package main

import (
	"math/rand"
	"time"
)

/**
 * Holds data for optimization
 */
type OptimizationData struct {
	A           float64
	Alpha       float64
	H           float64
	N           int
	MinBound    int
	MaxBound    int
	Precision   float64
	ThreadCount int
	Points      []Point
	FinalPoints []Point
	Duration 	int64
	Iterations  int
}

/**
 * 2D point
 */
type Point struct {
	X float64
	Y float64
}

/**
 * Gets data for optimization
 */
func getOptimizationData() []OptimizationData {
	data := make([]OptimizationData, 10)

	data[0] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 1,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[1] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 2,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[2] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 3,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[3] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 4,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[4] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 10,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[5] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 6,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[6] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 7,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[7] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 8,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[8] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 9,
		Points:      getRandomPoints(10, -10, 10),
	}

	data[9] = OptimizationData {
		A:           10,
		Alpha:       0.1,
		H:           1e-6,
		N:           10,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 10,
		Points:      getRandomPoints(10, -10, 10),
	}

	return data
}

/**
 * Generates the given number of random 2D points in the given bounds
 */
func getRandomPoints(count int, minBound, maxBound float64) []Point {
	points := make([]Point, count)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		points[i] = Point {X: randomFloat(minBound, maxBound), Y: randomFloat(minBound, maxBound)}
	}

	return points
}
