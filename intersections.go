package main

//Intersection ...
type Intersection struct {
	t      float64
	object Object
	index  int
}

//Intersections is a group of intersections
type Intersections struct {
	queue *PriorityQueue
}

//NewIntersections returns an Intersections
func NewIntersections(intersections []*Intersection) *Intersections {
	pq := PriorityQueue(intersections)
	pq.Init()

	return &Intersections{&pq}
}

//NewIntersection creates a new Intersection
func NewIntersection(t float64, object Object) *Intersection {
	return &Intersection{t, object, -1}
}

//Hit returns the object that the ray will Hit
func (xs *Intersections) Hit() *Intersection {
	if xs.queue.Empty() {
		return nil
	}
	top := xs.queue.Top()
	if top.t < 0 {
		return nil
	}
	return top
}

//Add adds a new intersection to the group
func (xs *Intersections) Add(intersection *Intersection) int {
	xs.queue.push(intersection)
	return xs.queue.Len()
}

//Count returns the size of the Intersections
func (xs *Intersections) Count() int {
	return xs.queue.Len()
}
