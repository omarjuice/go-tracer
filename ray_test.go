package main

import "testing"

func TestRayPosition(t *testing.T) {
	ray := NewRay(Point(2, 3, 4), Vector(1, 0, 0))

	results := []*Tuple{
		ray.Position(0),
		ray.Position(1),
		ray.Position(-1),
		ray.Position(2.5),
	}
	expected := []*Tuple{
		Point(2, 3, 4),
		Point(3, 3, 4),
		Point(1, 3, 4),
		Point(4.5, 3, 4),
	}
	for i := 0; i < 4; i++ {
		if !results[i].Equals(expected[i]) {
			t.Errorf("RayPosition: expected %v to be %v", results[i], expected[i])
		}
	}

}

func TestIntersectSphere(t *testing.T) {
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))

	s := NewSphere()

	xs := r.Intersect(s)

	if len(xs) != 2 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 2, len(xs))
	}

	expected := []float64{4.0, 6.0}

	for i, intersection := range xs {
		if !floatEqual(expected[i], intersection.t) {
			t.Errorf("IntersectSphere: expected %v to be %v", intersection.t, expected[i])
		}
	}

	r = NewRay(Point(0, 1, -5), Vector(0, 0, 1))
	xs = r.Intersect(s)

	if len(xs) != 2 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 2, len(xs))
	}
	expected = []float64{5.0, 5.0}

	for i, intersection := range xs {
		if !floatEqual(expected[i], intersection.t) {
			t.Errorf("IntersectSphere: expected %v to be %v", intersection.t, expected[i])
		}
	}

	r = NewRay(Point(0, 2, -5), Vector(0, 0, 1))
	xs = r.Intersect(s)

	if len(xs) != 0 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 0, len(xs))
	}

	r = NewRay(Point(0, 0, 5), Vector(0, 0, 1))
	xs = r.Intersect(s)

	if len(xs) != 2 {
		t.Errorf("IntersectSphere: expected number of intersections to be %v but got %v", 2, len(xs))
	}
	expected = []float64{-6.0, -4.0}

	for i, intersection := range xs {
		if !floatEqual(expected[i], intersection.t) {
			t.Errorf("IntersectSphere: expected %v to be %v", intersection.t, expected[i])
		}
	}

}
func TestHit(t *testing.T) {
	s := NewSphere()
	i1 := &Intersection{1, s}
	i2 := &Intersection{2, s}
	xs := NewIntersections(i1, i2)
	i := xs.Hit()
	if i != i1 {
		t.Errorf("Hit: expected %v to be %v", i, i1)
	}

	i1 = &Intersection{-1, s}
	i2 = &Intersection{2, s}
	xs = NewIntersections(i1, i2)
	i = xs.Hit()
	if i != i2 {
		t.Errorf("Hit: expected %v to be %v", i, i2)
	}

	i1 = &Intersection{-1, s}
	i2 = &Intersection{-2, s}
	xs = NewIntersections(i1, i2)
	i = xs.Hit()
	if i != nil {
		t.Errorf("Hit: expected %v to be %v", i, nil)
	}

	i1 = &Intersection{5, s}
	i2 = &Intersection{7, s}
	i3 := &Intersection{-3, s}
	i4 := &Intersection{2, s}
	xs = NewIntersections(i1, i2, i3, i4)
	i = xs.Hit()
	if i != i4 {
		t.Errorf("Hit: expected %v to be %v", i, i4)
	}

}
