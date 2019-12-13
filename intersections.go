package main

import (
	"sort"
)

//Intersection ...
type Intersection struct {
	t      float64
	object Shape
	index  int
}

//Intersections is a group of intersections
type Intersections []*Intersection

//NewIntersections returns an Intersections
func NewIntersections(intersections []*Intersection) Intersections {
	return Intersections(intersections)
}

//NewIntersection creates a new Intersection
func NewIntersection(t float64, object Shape) *Intersection {
	return &Intersection{t, object, -1}
}

//Hit returns the object that the ray will Hit
func (xs Intersections) Hit() *Intersection {

	sort.Slice(xs, func(i, j int) bool { return xs[i].t < xs[j].t })

	for _, i := range xs {
		if i.t >= 0.0 {
			return i
		}
	}
	return nil
}

//Add adds a new intersection to the group
func (xs *Intersections) Add(intersection *Intersection) int {
	*xs = append(*xs, intersection)
	return len(*xs)
}

//Count returns the size of the Intersections
func (xs *Intersections) Count() int {
	return len(*xs)
}
