package main

import (
	"math"
	"testing"
)

func TestWorldIntersections(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))

	xs := w.Intersect(r)
	count := xs.Count()
	if count != 4 {
		t.Errorf("WorldIntersections: %v should be %v", count, 4)
	}

	expected := []float64{4, 4.5, 5.5, 6}

	for _, v := range expected {
		top := xs.Hit()
		if top.t != v {
			t.Errorf("WorldIntersections: expected hit %v to be %v", top.t, v)
		}
		xs.queue.pop()
	}

}

func TestPrepareComputation(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	shape := NewSphere()
	i := &Intersection{4, shape, -1}
	comps := PrepareComputations(i, r)

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
	comps = PrepareComputations(i, r)

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
	comps := PrepareComputations(i, r)
	result := w.ShadeHit(comps, 10)
	expected := NewColor(0.38066, 0.47583, 0.2855)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}

	w.lights[0] = NewPointLight(Point(0, .25, 0), NewColor(1, 1, 1))
	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape = w.objects[1]
	i = NewIntersection(0.5, shape)
	comps = PrepareComputations(i, r)
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

	comps = PrepareComputations(i, r)
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
	comps := PrepareComputations(i, r)
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
	comps := PrepareComputations(i, r)

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
	comps = PrepareComputations(i, r)

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
	comps := PrepareComputations(i, r)

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
