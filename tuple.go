package main

import "math"

//Tuple is a tuple of four floating points
type Tuple struct {
	x, y, z, w float64
}

//Point creates a 3D Point
func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

// Vector creates a 3D movement vector
func Vector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

//IsPoint tells whether it is a Point or not
func (t *Tuple) IsPoint() bool {
	return t.w == 1.0
}

//IsVector tells whether it is a vector or not
func (t *Tuple) IsVector() bool {
	return t.w == 0.0
}

//Equals returns whether a t is equal to ano
func (t *Tuple) Equals(o *Tuple) bool {

	return FloatEqual(t.x, o.x) && FloatEqual(t.y, o.y) && FloatEqual(t.z, o.z) && FloatEqual(t.w, o.w)
}

//Add adds two ts together
func (t *Tuple) Add(o *Tuple) *Tuple {
	return &Tuple{
		t.x + o.x,
		t.y + o.y,
		t.z + o.z,
		t.w + o.w,
	}
}

//Sub subtracts two ts
func (t *Tuple) Sub(o *Tuple) *Tuple {
	return &Tuple{
		t.x - o.x,
		t.y - o.y,
		t.z - o.z,
		t.w - o.w,
	}

}

//Mul multiplies a t by a scalar
func (t *Tuple) Mul(scalar float64) *Tuple {
	return &Tuple{
		t.x * scalar,
		t.y * scalar,
		t.z * scalar,
		t.w * scalar,
	}

}

//Div divides a t by a scalar
func (t *Tuple) Div(scalar float64) *Tuple {
	return &Tuple{
		t.x / scalar,
		t.y / scalar,
		t.z / scalar,
		t.w / scalar,
	}

}

// Negate negates the t
func (t *Tuple) Negate() *Tuple {
	return &Tuple{-t.x, -t.y, -t.z, -t.w}
}

func square(v float64) float64 {
	return math.Pow(v, 2.0)
}

//Magnitude returns the magnitude of a vector
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(square(t.x) +
		square(t.y) +
		square(t.z) +
		square(t.w))
}

//Normalize normalizes a vector
func (t *Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	if mag == 0.0 {
		return t
	}
	return Vector(t.x/mag, t.y/mag, t.z/mag)

}

//Dot calculates the dot product of two vectors
func (t *Tuple) Dot(o *Tuple) float64 {
	return t.x*o.x + t.y*o.y + t.z*o.z + t.w*o.w
}

//Cross calculates the cross product of two vectors
func (t *Tuple) Cross(o *Tuple) *Tuple {
	return Vector(
		t.y*o.z-t.z*o.y,
		t.z*o.x-t.x*o.z,
		t.x*o.y-t.y*o.x,
	)
}

//String converts a tuple to a string
func (t *Tuple) String() string {
	start := "("
	if t.IsPoint() {
		start = "p" + start
	} else {
		start = "v" + start
	}
	return start + FloatToString(t.x) + "," + FloatToString(t.y) + "," + FloatToString(t.z) + ")"
}
