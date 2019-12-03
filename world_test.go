package main

import (
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
	result := w.ShadeHit(comps)
	expected := NewColor(0.38066, 0.47583, 0.2855)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}

	w.lights[0] = NewPointLight(Point(0, .25, 0), NewColor(1, 1, 1))
	r = NewRay(Point(0, 0, 0), Vector(0, 0, 1))
	shape = w.objects[1]
	i = NewIntersection(0.5, shape)
	comps = PrepareComputations(i, r)
	result = w.ShadeHit(comps)
	expected = NewColor(0.90498, 0.90498, 0.90498)

	if !result.Equals(expected) {
		t.Errorf("ShadeHit: expected %v to be %v", result, expected)
	}
}

func TestWorldColorAt(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(Point(0, 0, -5), Vector(0, 1, 0))
	result := w.ColorAt(r)
	expected := Black

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (no hit): expected %v to be %v", result, expected)
	}

	r = NewRay(r.origin, Vector(0, 0, 1))
	result = w.ColorAt(r)
	expected = NewColor(0.38066, 0.47583, 0.2855)

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (hit): expected %v to be %v", result, expected)
	}

	outer := w.objects[0]
	outer.Material().ambient = 1
	inner := w.objects[1]

	inner.Material().ambient = 1

	r = NewRay(Point(0, 0, .75), Vector(0, 0, -1))

	result = w.ColorAt(r)
	expected = inner.Material().color

	if !result.Equals(expected) {
		t.Errorf("WorldColorAt (hit inner): expected %v to be %v", result, expected)
	}

}
