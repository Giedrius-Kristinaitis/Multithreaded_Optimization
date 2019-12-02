package main

/**
 * Gets a deep copy of the given point array
 *
 * (Yes, I just reinvented a wheel)
 */
func deepCopyPointArray(points []Point) []Point {
	deepCopy := make([]Point, len(points))

	for index, point := range points {
		deepCopy[index] = Point {X: point.X, Y: point.Y}
	}

	return deepCopy
}
