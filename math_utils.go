package main

import (
	"math"
	"math/rand"
)

/**
 * A gradient value passed through a channel
 */
type GradientMessage struct {
	Index int
	X     float64
	Y     float64
}

/**
 * Calculates distance between two points
 */
func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

/**
 * Generates a random float in the given bounds
 */
func randomFloat(minBound, maxBound float64) float64 {
	return minBound + rand.Float64()*math.Abs(maxBound-minBound)
}

/**
 * Gets the sum of the distances between all points
 */
func getPointDistanceSum(points []Point, a float64) float64 {
	sum := 0.0

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			sum += math.Pow(distance(points[i].X, points[i].Y, points[j].X, points[j].Y) - a, 2)
		}
	}

	return sum / float64(len(points))
}

/**
 * Gets all point gradients
 */
func getAllPointGradients(gradientChannel chan GradientMessage, allPoints []Point, currentDistanceSum float64, h float64, a float64) []Point {
	gradients := make([]Point, len(allPoints))

	for i := 0; i < len(allPoints); i++ {
		go getPointGradient(gradientChannel, i, allPoints, currentDistanceSum, h, a)
	}

	for i := 0; i < len(allPoints); i++ {
		gradientMessage := <- gradientChannel
		gradients[gradientMessage.Index] = Point {X: gradientMessage.X, Y: gradientMessage.Y}
	}

	return gradients
}

/**
 * Gets the gradient vector
 */
func getPointGradient(gradientChannel chan GradientMessage, pointIndex int, allPoints []Point, currentDistanceSum float64, h float64, a float64) {
	newPoints := deepCopyPointArray(allPoints)

	newPoints[pointIndex].X += h

	distanceSumChangedX := getPointDistanceSum(newPoints, a);
	newPoints[pointIndex].X -= h
	newPoints[pointIndex].Y += h

	distanceSumChangedY := getPointDistanceSum(newPoints, a)

	gradientX := (distanceSumChangedX - currentDistanceSum) / h
	gradientY := (distanceSumChangedY - currentDistanceSum) / h

	gradientChannel <- GradientMessage {Index: pointIndex, X: gradientX, Y: gradientY}
}

/**
 * Normalizes a gradient vector
 */
func normalizeGradient(gradient []float64) []float64 {
	norm := vectorNorm(gradient)

	for index, value := range gradient {
		gradient[index] = value / norm
	}

	return gradient
}

/**
 * Gets the norm of a vector
 */
func vectorNorm(gradient []float64) float64 {
	sum := 0.0

	for _, value := range gradient {
		sum += math.Pow(value, 2)
	}

	return math.Sqrt(sum)
}

/**
 * Changes the coordinates of each point to normalize distances between points
 */
func changePointCoordinates(points []Point, gradients []Point, alpha float64) []Point {
	newPoints := deepCopyPointArray(points)

	for index := 0; index < len(points); index++ {
		newPoints[index].X -= alpha * gradients[index].X
		newPoints[index].Y -= alpha * gradients[index].Y
	}

	return newPoints
}
