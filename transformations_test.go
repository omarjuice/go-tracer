package main

import (
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	transform := Translation(5, -3, 2)
	p := Point(-3, 4, 5)

	result := transform.MulTuple(p)

	expected := Point(2, 1, 7)

	if !result.Equals(expected) {
		t.Errorf("Translation:Point expected %v to equal %v", result, expected)
	}

	inv := transform.Inverse()

	result = inv.MulTuple(p)
	expected = Point(-8, 7, 3)

	if !result.Equals(expected) {
		t.Errorf("Translation:Point expected %v to equal %v", result, expected)
	}

	v := Vector(-3, 4, 5)

	if !transform.MulTuple(v).Equals(v) {
		t.Errorf("Translation:Vector vector was changed by translation.")
	}

}

func TestScaling(t *testing.T) {
	transform := Scaling(2, 3, 4)
	p := Point(-4, 6, 8)

	result := transform.MulTuple(p)
	expected := Point(-8, 18, 32)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Point expected %v to equal %v", result, expected)
	}

	v := Vector(-4, 6, 8)

	result = transform.MulTuple(v)
	expected = Vector(-8, 18, 32)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Vector expected %v to equal %v", result, expected)
	}

	inv := transform.Inverse()
	result = inv.MulTuple(v)
	expected = Vector(-2, 2, 2)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Vector (inverse) expected %v to equal %v", result, expected)
	}

	transform = Scaling(-1, 1, 1)
	p = Point(2, 3, 4)
	result = transform.MulTuple(p)
	expected = Point(-2, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Scaling:Point (reflection) expected %v to equal %v", result, expected)
	}
}

func TestRotationX(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationX(π / 4)

	result := halfQuarter.MulTuple(p)

	expected := Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !result.Equals(expected) {

		t.Errorf("RotationX: expected %v to be %v", result, expected)
	}

	fullQuarter := RotationX(π / 2)

	expected = Point(0, 0, 1)

	result = fullQuarter.MulTuple(p)

	if !result.Equals(expected) {
		t.Errorf("RotationX: expected %v to be %v", result, expected)
	}

	invHalfQuarter := halfQuarter.Inverse()

	expected = Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	result = invHalfQuarter.MulTuple(p)

	if !result.Equals(expected) {
		t.Errorf("RotationX: expected %v to be %v", result, expected)
	}
}
func TestRotationY(t *testing.T) {
	p := Point(0, 0, 1)
	halfQuarter := RotationY(π / 4)

	result := halfQuarter.MulTuple(p)

	expected := Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}

	fullQuarter := RotationY(π / 2)

	result = fullQuarter.MulTuple(p)

	expected = Point(1, 0, 0)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}
}

func TestRotationZ(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationZ(π / 4)

	result := halfQuarter.MulTuple(p)

	expected := Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}

	fullQuarter := RotationZ(π / 2)

	result = fullQuarter.MulTuple(p)

	expected = Point(-1, 0, 0)

	if !result.Equals(expected) {
		t.Errorf("RotationY expected %v to be %v", result, expected)
	}

}

func TestShearing(t *testing.T) {
	p := Point(2, 3, 4)

	transform := Shearing(1, 0, 0, 0, 0, 0)
	result := transform.MulTuple(p)
	expected := Point(5, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: xy expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 1, 0, 0, 0, 0)
	result = transform.MulTuple(p)
	expected = Point(6, 3, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: xz expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 1, 0, 0, 0)
	result = transform.MulTuple(p)
	expected = Point(2, 5, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: yx expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 0, 1, 0, 0)
	result = transform.MulTuple(p)
	expected = Point(2, 7, 4)

	if !result.Equals(expected) {
		t.Errorf("Shearing: yz expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 0, 0, 1, 0)
	result = transform.MulTuple(p)
	expected = Point(2, 3, 6)

	if !result.Equals(expected) {
		t.Errorf("Shearing: zx expected %v to be %v", result, expected)
	}

	transform = Shearing(0, 0, 0, 0, 0, 1)
	result = transform.MulTuple(p)
	expected = Point(2, 3, 7)

	if !result.Equals(expected) {
		t.Errorf("Shearing: zy expected %v to be %v", result, expected)
	}
}

func TestChainTransformations(t *testing.T) {
	p := Point(1, 0, 1)
	A := RotationX(π / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)

	result := p.Transform(A, B, C)

	expected := Point(15, 0, 7)

	if !result.Equals(expected) {
		t.Errorf("ChainTransformations: expected %v to be %v", result, expected)
	}

}

func TestViewTransform(t *testing.T) {
	from := Point(0, 0, 0)
	to := Point(0, 0, -1)
	up := Vector(0, 1, 0)

	result := ViewTransform(from, to, up)
	expected := IdentityMatrix

	if !result.Equals(expected) {
		t.Errorf("ViewTransform(default): expected %v to equal %v", result, expected)
	}

	from = Point(0, 0, 0)
	to = Point(0, 0, 1)
	up = Vector(0, 1, 0)
	result = ViewTransform(from, to, up)
	expected = Scaling(-1, 1, -1)

	if !result.Equals(expected) {
		t.Errorf("ViewTransform(positive z): expected %v to equal %v", result, expected)
	}

	from = Point(0, 0, 8)
	to = Point(0, 0, 0)
	up = Vector(0, 1, 0)
	result = ViewTransform(from, to, up)
	expected = Translation(0, 0, -8)

	if !result.Equals(expected) {
		t.Errorf("ViewTransform(moves world): expected %v to equal %v", result, expected)
	}

	from = Point(1, 3, 2)
	to = Point(4, -2, 8)
	up = Vector(1, 1, 0)
	result = ViewTransform(from, to, up)
	expected = [][]float64{
		[]float64{-0.507093, 0.507093, 0.676123, -2.366432},
		[]float64{0.767716, 0.606092, 0.121218, -2.828427},
		[]float64{-0.358569, 0.597614, -0.717137, 0},
		[]float64{0, 0, 0, 1},
	}
	if !result.Equals(expected) {
		t.Errorf("ViewTransform(arbitary): expected %v to equal %v", result, expected)
	}

}
