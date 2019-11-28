package main

import "sort"

//Intersection ...
type Intersection struct {
	t      float64
	object Object
}

//Intersections is a group of intersections
type Intersections struct {
	list []*Intersection
}

//NewIntersections returns an Intersections
func NewIntersections(intersections ...*Intersection) *Intersections {
	return &Intersections{intersections}
}

//Hit returns the object that the ray will Hit
func (xs *Intersections) Hit() *Intersection {

	sort.Slice(xs.list, func(i, j int) bool { return xs.list[i].t < xs.list[j].t })

	for _, intersection := range xs.list {
		if intersection.t >= 0 {
			return intersection
		}
	}
	return nil
}

//Add adds a new intersection to the group
func (xs *Intersections) Add(intersection *Intersection) int {
	xs.list = append(xs.list, intersection)
	return len(xs.list)
}
