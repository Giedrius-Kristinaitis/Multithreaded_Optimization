package main

import (
	"fmt"
	"time"
)

/**
 * Some constants
 */
const MAX_ITERATIONS = 10000

/**
 * Entry point of the program
 */
func main() {
	data := getOptimizationData()

	for _, dataSet := range data {
		optimize(dataSet)
	}
}

/**
 * Performs optimization with the given optimization data and prints the results
 */
func optimize(data OptimizationData) {
	startTime := time.Now()

	data = optimizePoints(data)

	data.Duration = time.Since(startTime).Milliseconds()

	printResults(data)
}

/**
 * Performs optimization with the given optimization data
 */
func optimizePoints(data OptimizationData) OptimizationData {
	data.FinalPoints = deepCopyPointArray(data.Points)
	gradientChannel := make(chan GradientMessage)
	iteration := 0
	currentDistanceSum := getPointDistanceSum(data.FinalPoints, data.A)

	for iteration <= MAX_ITERATIONS && data.Alpha > data.Precision {
		iteration++

		allGradients := getAllPointGradients(gradientChannel, data.FinalPoints, currentDistanceSum, data.H, data.A)
		changedPoints := changePointCoordinates(data.FinalPoints, allGradients, data.Alpha)

		changedDistanceSum := getPointDistanceSum(changedPoints, data.A)

		if changedDistanceSum < currentDistanceSum {
			data.FinalPoints = changedPoints
			currentDistanceSum = changedDistanceSum
		} else {
			data.Alpha /= 2
		}
	}

	data.Iterations = iteration

	return data
}

/**
 * Prints optimization results
 */
func printResults(data OptimizationData) {
	/*for i, point := range data.Points {
		fmt.Printf("Start point %d X %f\n", i, point.X)
		fmt.Printf("Start point %d Y %f\n", i, point.X)
	}

	for i, point := range data.FinalPoints {
		fmt.Printf("End point %d X %f\n", i, point.X)
		fmt.Printf("End point %d Y %f\n", i, point.X)
	}*/

	fmt.Printf("Duration: %d ms\n", data.Duration)
	fmt.Printf("Iterations: %d\n\n", data.Iterations)
}
