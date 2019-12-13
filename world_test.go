package main

import (
	"math"
	"testing"
)

func TestWorldIntersections(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))

	xs := w.Intersect(r)
	xs.Hit()
	count := xs.Count()
	if count != 4 {
		t.Errorf("WorldIntersections: %v should be %v", count, 4)
	}

	expected := []float64{4, 4.5, 5.5, 6}

	for i, v := range expected {
		top := xs[i]
		if top.t != v {
			t.Errorf("WorldIntersections: expected hit %v to be %v", top.t, v)
		}
	}

}

func TestPrepareComputation(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	i := &Intersection{4, shape, -1}
	comps := PrepareComputations(i, r, NewIntersections([]*Intersection{i}))

	if !floatEqual(comps.t, i.t) {
		t.Errorf("PrepareComputations failed")
	}

	if comps.object != i.object {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.point.Equals(Point(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.eyev.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.normalv.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if comps.inside {
		t.Errorf("PrepareComputations failed")
	}

	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape = NewSphere()
	i = &Intersection{1, shape, -1}
	comps = PrepareComputations(i, r, NewIntersections([]*Intersection{i}))

	if !comps.point.Equals(Point(0, 0, 1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.eyev.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.normalv.Equals(Vector(0, 0, -1)) {
		t.Errorf("PrepareComputations failed")
	}
	if !comps.inside {
		t.Errorf("PrepareComputations failed")
	}

}

func TestShadeHit(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := w.objects[0]
	i := NewIntersection(4, shape)
	comps := PrepareComputations(i, r, NewIntersections([]*Intersection{i}))
	result := w.ShadeHit(comps, 10)
	expected := NewColor(0.38066, 0.47583, 0.2855)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}

	w.lights[0] = NewPointLight(Point(0, .25, 0), NewColor(1, 1, 1))
	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape = w.objects[1]
	i = NewIntersection(0.5, shape)
	comps = PrepareComputations(i, r, NewIntersections([]*Intersection{i}))
	result = w.ShadeHit(comps, 10)
	expected = NewColor(0.90498, 0.90498, 0.90498)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}

	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, 10))

	w = NewWorld(
		[]*PointLight{NewPointLight(Point(0, 0, -10), NewColor(1, 1, 1))},
		[]Shape{s1, s2},
	)
	r = NewRay(Point(0, 0, 5), Vector(0, 0, 1))
	i = NewIntersection(4, s2)

	comps = PrepareComputations(i, r, NewIntersections([]*Intersection{i}))
	result = w.ShadeHit(comps, 10)
	expected = NewColor(0.1, 0.1, 0.1)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}
}

func TestWorldColorAt(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 1, 0))
	result := w.ColorAt(r, 10)
	expected := Black

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (no hit): expected %v to be %v", result, expected)
	}

	r = NewRay(r.origin, Vector(0, 0, 1))
	result = w.ColorAt(r, 10)
	expected = NewColor(0.38066, 0.47583, 0.2855)

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (hit): expected %v to be %v", result, expected)
	}

	outer := w.objects[0]
	outer.Material().ambient = 1
	inner := w.objects[1]

	inner.Material().ambient = 1

	r = NewRay(Point(0, 0, .75), Vector(0, 0, -1))

	result = w.ColorAt(r, 10)
	expected = inner.Material().color

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (hit inner): expected %v to be %v", result, expected)
	}

}

func TestIsShadowed(t *testing.T) {
	w := DefaultWorld()

	p := Point(0, 10, 0)
	if w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: expected no shadow when nothing is collinear point and light")
	}

	p = Point(10, -10, 10)
	if !w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: expected object between point and light to create shadow")
	}

	p = Point(-20, 20, -20)
	if w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: There should be no shadow when an object is behind the light")
	}

	p = Point(-2, 2, -2)
	if w.IsShadowed(p, 0) {
		t.Errorf("IsShadowed: There is no shadow when an object is behind the point ")
	}

}

func TestComputeReflect(t *testing.T) {
	shape := NewPlane()
	r := NewRay(Point(0, 1, -1), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), shape)
	comps := PrepareComputations(i, r, NewIntersections([]*Intersection{i}))
	expected := Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !comps.reflectv.Equals(expected) {
		t.Errorf("PrepareComputationsWithReflect: expected %v to be %v", comps.reflectv, expected)
	}

}

func TestWorldReflect(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))

	shape := w.objects[1]

	shape.Material().ambient = 1
	i := NewIntersection(1, shape)
	comps := PrepareComputations(i, r, NewIntersections([]*Intersection{i}))

	color := w.ReflectedColor(comps, 10)
	if !color.Equals(Black) {
		t.Errorf("WorldReflect(non-reflective): expected %v to be %v", color, Black)
	}

	shape = NewPlane()
	shape.Material().reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.objects = append(w.objects, shape)
	r = NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i = NewIntersection(math.Sqrt(2), shape)
	comps = PrepareComputations(i, r, NewIntersections([]*Intersection{i}))

	color = w.ReflectedColor(comps, 10)
	expected := NewColor(0.19033, 0.237915, 0.142749)

	if !color.Equals(expected) {
		t.Errorf("WorldReflect(reflective): expected %v to be %v", color, expected)
	}
}

func TestShadeHitWithReflect(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.Material().reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.objects = append(w.objects, shape)
	r := NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), shape)
	comps := PrepareComputations(i, r, NewIntersections([]*Intersection{i}))

	color := w.ShadeHit(comps, 10)
	expected := NewColor(0.876758, 0.924341, 0.829175)

	if !color.Equals(expected) {
		t.Errorf("WorldReflect(reflective): expected %v to be %v", color, expected)
	}

}

func TestInfiniteReflection(t *testing.T) {
	light := NewPointLight(Point(0, 0, 0), NewColor(1, 1, 1))

	lower := NewPlane()
	lower.Material().reflective = 1
	lower.SetTransform(Translation(0, -1, 0))

	upper := NewPlane()
	upper.Material().reflective = 1
	upper.SetTransform(Translation(0, 1, 0))

	w := NewWorld([]*PointLight{light}, []Shape{lower, upper})

	r := NewRay(Point(0, 0, 0), Vector(0, 1, 0))

	w.ColorAt(r, 10)

}

func TestPrepareComputationWithRefraction(t *testing.T) {
	A := GlassSphere()
	A.SetTransform(Scaling(2, 2, 2))
	A.Material().refractiveIndex = 1.5

	B := GlassSphere()
	B.SetTransform(Translation(0, 0, -.25))
	B.Material().refractiveIndex = 2.0

	C := GlassSphere()
	C.SetTransform(Translation(0, 0, 0.25))
	C.Material().refractiveIndex = 2.5

	r := NewRay(Point(0, 0, -4), Vector(0, 0, 1))
	xs := NewIntersections([]*Intersection{NewIntersection(2, A), NewIntersection(2.75, B), NewIntersection(3.25, C), NewIntersection(4.75, B), NewIntersection(5.25, C), NewIntersection(6, A)})

	examples := map[int][2]float64{
		0: [2]float64{1.0, 1.5},
		1: [2]float64{1.5, 2.0},
		2: [2]float64{2.0, 2.5},
		3: [2]float64{2.5, 2.5},
		4: [2]float64{2.5, 1.5},
		5: [2]float64{1.5, 1.0},
	}

	for idx, N := range examples {
		comps := PrepareComputations(xs[idx], r, xs)
		n1, n2 := N[0], N[1]
		if !floatEqual(comps.n1, n1) || !floatEqual(comps.n2, n2) {
			t.Errorf("PrepareComputationWithRefraction: Expected %v,%v to be %v,%v", comps.n1, comps.n2, n1, n2)
		}
	}

	r = NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := GlassSphere()
	shape.SetTransform(Translation(0, 0, 1))
	i := NewIntersection(5, shape)
	xs = NewIntersections([]*Intersection{i})
	comps := PrepareComputations(i, r, xs)

	if !(comps.underPoint.z > EPSILON/2.0 && comps.point.z < comps.underPoint.z) {
		t.Errorf("PrepareComputationWithRefraction: underPoint %v not valid", comps.underPoint)
	}

}

func TestWorldRefractedColor(t *testing.T) {
	w := DefaultWorld()
	shape := w.objects[0]
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	xs := NewIntersections([]*Intersection{NewIntersection(4, shape), NewIntersection(6, shape)})

	comps := PrepareComputations(xs[0], r, xs)
	c := w.RefractedColor(comps, 5)

	if !c.Equals(Black) {
		t.Errorf("WorldRefractedColor(opaque): expected %v to be %v", c, Black)
	}

	shape.Material().transparency = 1.0
	shape.Material().refractiveIndex = 1.5

	comps = PrepareComputations(xs[0], r, xs)

	c = w.RefractedColor(comps, 0)
	if !c.Equals(Black) {
		t.Errorf("WorldRefractedColor(no remaining recursion): expected %v to be %v", c, Black)
	}

	r = NewRay(Point(0, 0, math.Sqrt(2)/2), Vector(0, 1, 0))
	xs = NewIntersections([]*Intersection{NewIntersection(-math.Sqrt(2)/2, shape), NewIntersection(math.Sqrt(2)/2, shape)})

	comps = PrepareComputations(xs[1], r, xs)

	c = w.RefractedColor(comps, 5)

	if !c.Equals(Black) {
		t.Errorf("WorldRefractedColor(total internal reflection): expected %v to be %v", c, Black)
	}

	w = DefaultWorld()
	A := w.objects[0]

	A.Material().ambient = 1.0
	A.Material().pattern = NewPattern([][]*Color{[]*Color{}}, func(colors []*Color, point *Tuple) *Color { return NewColor(point.x, point.y, point.z) })

	B := w.objects[1]
	B.Material().transparency = 1.0
	B.Material().refractiveIndex = 1.5

	r = NewRay(Point(0, 0, 0.1), Vector(0, 1, 0))
	xs = NewIntersections([]*Intersection{NewIntersection(-.9899, A), NewIntersection(-.4899, B), NewIntersection(.4899, B), NewIntersection(.9899, A)})

	comps = PrepareComputations(xs[2], r, xs)

	c = w.RefractedColor(comps, 5)

	expected := NewColor(0.000000, 0.998875, 0.047219)

	if !c.Equals(expected) {
		t.Errorf("WorldRefractedColor(actual refraction): expected %v to be %v", c, expected)
	}

}

func TestWorldShadeHitWithRefraction(t *testing.T) {
	w := DefaultWorld()
	floor := NewPlane()
	floor.SetTransform(Translation(0, -1, 0))
	floor.Material().transparency = 0.5
	floor.Material().refractiveIndex = 1.5
	w.objects = append(w.objects, floor)

	ball := NewSphere()
	ball.Material().color = NewColor(1, 0, 0)
	ball.Material().ambient = 0.5
	ball.SetTransform(Translation(0, -3.5, -0.5))
	w.objects = append(w.objects, ball)

	r := NewRay(Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := NewIntersections([]*Intersection{NewIntersection(math.Sqrt(2), floor)})

	comps := PrepareComputations(xs[0], r, xs)
	color := w.ShadeHit(comps, 5)
	expected := NewColor(0.936425, 0.686425, 0.686425)

	if !color.Equals(expected) {
		t.Errorf("WorldShadeHitWithRefraction(no reflection): expected %v to be %v", color, expected)
	}

	floor.Material().reflective = 0.5

	comps = PrepareComputations(xs[0], r, xs)
	color = w.ShadeHit(comps, 5)
	expected = NewColor(0.933915, 0.696434, 0.692430)
	if !color.Equals(expected) {
		t.Errorf("WorldShadeHitWithRefraction(with reflection): expected %v to be %v", color, expected)
	}
}

func TestComputationSchlick(t *testing.T) {
	shape := GlassSphere()
	r := NewRay(Point(0, 0, math.Sqrt(2)/2), Vector(0, 1, 0))
	xs := NewIntersections([]*Intersection{NewIntersection(-math.Sqrt(2)/2, shape), NewIntersection(math.Sqrt(2)/2, shape)})
	xs.Hit()
	comps := PrepareComputations(xs[1], r, xs)
	reflectance := comps.Schlick()
	expected := 1.0

	if !floatEqual(reflectance, expected) {
		t.Errorf("Schlick(internal reflection): expected %v to be %v", reflectance, expected)
	}

	r = NewRay(Point(0, 0, 0), Vector(0, 1, 0))
	xs = NewIntersections([]*Intersection{NewIntersection(-1, shape), NewIntersection(1, shape)})
	comps = PrepareComputations(xs[1], r, xs)
	reflectance = comps.Schlick()
	expected = 0.04

	if !floatEqual(reflectance, expected) {
		t.Errorf("Schlick(perpendicular viewing): expected %v to be %v", reflectance, expected)
	}

	r = NewRay(Point(0, 0.99, -2), Vector(0, 0, 1))
	xs = NewIntersections([]*Intersection{NewIntersection(1.8589, shape)})
	comps = PrepareComputations(xs[0], r, xs)
	reflectance = comps.Schlick()
	expected = 0.48873

	if !floatEqual(reflectance, expected) {
		t.Errorf("Schlick(small angle, n2 > n1): expected %v to be %v", reflectance, expected)
	}

}
