package main

import "strconv"

//EPSILON is equivalence tolerance for float value comparison
const EPSILON = 0.00001

//floatEqual determines if two floats are equal within a tolerance
func floatEqual(a, b float64) bool {
	return abs(a-b) < EPSILON
}
func floatToUint8String(f float64) string {
	if f < 0.0 {
		return "0"
	}
	f *= 256.0
	if f > 255.0 {
		return "255"
	}
	return strconv.Itoa(int(f))
}

//Abs returns absolute value
func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n

}

//Sum adds a sequence together
func sum(nums ...float64) float64 {
	result := 0.0

	for _, v := range nums {
		result += v
	}
	return result
}

//FloatToString converts a float to a String
func floatToString(n float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(n, 'f', 6, 64)[:3]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
