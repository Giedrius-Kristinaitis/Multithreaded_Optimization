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
 * Data passed into get point gradient method to avoid using shared memory
 */
type GradientDataMessage struct {
	PointIndex	int
	EndIndex	int
	Points		[]Point
	CurrentSum	float64
	A 			float64
	H 			float64
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
func getAllPointGradients(gradientChannel chan GradientMessage, gradientDataChannel chan GradientDataMessage, threadCount int, allPoints []Point, currentDistanceSum float64, h float64, a float64) []Point {
	gradients := make([]Point, len(allPoints))

	for i := 0; i < threadCount; i++ {
		go getPointGradient(gradientChannel, gradientDataChannel)
	}

	for i := 0; i < threadCount; i++ {
		startIndex :=(len(allPoints) / threadCount) * i
		endIndex := startIndex + (len(allPoints) / threadCount)

		if len(allPoints) - endIndex < len(allPoints) / threadCount {
			endIndex = len(allPoints) - 1
		}

		gradientDataChannel <- GradientDataMessage {
			PointIndex: startIndex,
			EndIndex: endIndex,
			Points: allPoints,
			CurrentSum: currentDistanceSum,
			A: a,
			H: h,
		}
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
func getPointGradient(gradientChannel chan GradientMessage, gradientDataChannel chan GradientDataMessage) {
	gradientDataMessage := <- gradientDataChannel

	for i := gradientDataMessage.PointIndex; i <= gradientDataMessage.EndIndex; i++ {
		newPoints := deepCopyPointArray(gradientDataMessage.Points)

		newPoints[i].X += gradientDataMessage.H

		distanceSumChangedX := getPointDistanceSum(newPoints, gradientDataMessage.A);
		newPoints[i].X -= gradientDataMessage.H
		newPoints[i].Y += gradientDataMessage.H

		distanceSumChangedY := getPointDistanceSum(newPoints, gradientDataMessage.A)

		gradientX := (distanceSumChangedX - gradientDataMessage.CurrentSum) / gradientDataMessage.H
		gradientY := (distanceSumChangedY - gradientDataMessage.CurrentSum) / gradientDataMessage.H

		gradientChannel <- GradientMessage{Index: i, X: gradientX, Y: gradientY}
	}
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
		gradient := make([]float64, 2)

		gradient[0] = gradients[index].X
		gradient[1] = gradients[index].Y

		gradient = normalizeGradient(gradient)

		newPoints[index].X -= alpha * gradient[0]
		newPoints[index].Y -= alpha * gradient[1]
	}

	return newPoints
}
