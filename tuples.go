package main

import "math"

//Tuple is a tuple of four floating points
type Tuple [4]float64

// X,Y,Z,W are respective indices of Tuples
const (
	X = iota
	Y
	Z
	W
)

//Point creates a 3D Point
func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// Vector creates a 3D movement vector
func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

//IsPoint tells whether it is a Point or not
func (tuple *Tuple) IsPoint() bool {
	return tuple[W] == 1.0
}

//IsVector tells whether it is a vector or not
func (tuple *Tuple) IsVector() bool {
	return tuple[W] == 0.0
}

//Equals returns whether a tuple is equal to another
func (tuple *Tuple) Equals(other Tuple) bool {
	for i, v := range tuple {
		if !FloatEqual(v, other[i]) {
			return false
		}
	}
	return true
}

//Add adds two tuples together
func (tuple *Tuple) Add(other Tuple) Tuple {
	result := Tuple{}
	for i := 0; i < 4; i++ {
		result[i] = tuple[i] + other[i]
	}
	return result
}

//Sub subtracts two tuples
func (tuple *Tuple) Sub(other Tuple) Tuple {
	result := Tuple{}
	for i := 0; i < 4; i++ {
		result[i] = tuple[i] - other[i]
	}
	return result
}

//Mul multiplies a tuple by a scalar
func (tuple *Tuple) Mul(scalar float64) Tuple {
	result := Tuple{}

	for i, v := range tuple {
		result[i] = v * scalar
	}
	return result
}

//Div divides a tuple by a scalar
func (tuple *Tuple) Div(scalar float64) Tuple {
	result := Tuple{}

	for i, v := range tuple {
		result[i] = v / scalar
	}
	return result
}

// Negate negates the tuple
func (tuple *Tuple) Negate() Tuple {
	result := Tuple{}

	for i, v := range tuple {
		result[i] = -v
	}
	return result
}

//Magnitude returns the magnitude of a vector
func (tuple *Tuple) Magnitude() float64 {
	sumSquares := 0.0

	for i := 0; i < 3; i++ {
		sumSquares += math.Pow(tuple[i], 2.0)
	}
	return math.Sqrt(sumSquares)
}

//Normalize normalizes a vector
func (tuple *Tuple) Normalize() Tuple {
	mag := tuple.Magnitude()
	if mag == 0.0 {
		return *tuple
	}
	result := Vector(0, 0, 0)

	for i := 0; i < 3; i++ {
		result[i] = tuple[i] / mag
	}
	return result
}

//Dot calculates the dot product of two vectors
func (tuple *Tuple) Dot(other Tuple) float64 {
	result := 0.0
	for i, v := range tuple {
		result += v * other[i]
	}

	return result
}

//Cross calculates the cross product of two vectors
func (tuple *Tuple) Cross(other Tuple) Tuple {
	return Vector(
		tuple[Y]*other[Z]-tuple[Z]*other[Y],
		tuple[Z]*other[X]-tuple[X]*other[Z],
		tuple[X]*other[Y]-tuple[Y]*other[X],
	)
}
