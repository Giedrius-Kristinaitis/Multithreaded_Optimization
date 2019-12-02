package main

/**
 * Holds data for optimization
 */
type OptimizationData struct {
	A           int
	Alpha       float64
	H           float64
	N           int
	MinBound    int
	MaxBound    int
	Precision   float64
	ThreadCount int
	Points      []Point
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
		N:           5,
		MinBound:    -10,
		MaxBound:    10,
		Precision:   1e-6,
		ThreadCount: 4,
		Points:      getRandomPoints(5, -10, 10),
	}

	return data
}

/**
 * Generates the given number of random 2D points in the given bounds
 */
func getRandomPoints(count int, minBound, maxBound float64) []Point {
	points := make([]Point, count)

	for i := 0; i < count; i++ {
		points[i] = Point {X: randomFloat(minBound, maxBound), Y: randomFloat(minBound, maxBound)}
	}

	return points
}
