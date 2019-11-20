package main

//EPSILON is equivalence tolerance for float value comparison
const EPSILON = 0.00001

//FloatEqual determines if two floats are equal within a tolerance
func FloatEqual(a, b float64) bool {
	return Abs(a-b) < EPSILON
}

//Abs returns absolute value
func Abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n

}

//Sum adds a sequence together
func Sum(nums ...float64) float64 {
	result := 0.0

	for _, v := range nums {
		result += v
	}
	return result
}
