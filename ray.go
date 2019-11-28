package main

//Ray is a ray casting primitive
type Ray struct {
	origin, direction *Tuple
}

//NewRay creates a new ray
func NewRay(origin, direction *Tuple) *Ray {
	return &Ray{origin, direction}
}

//Position calculates the point at the given distance t along the ray
func (ray *Ray) Position(t float64) *Tuple {
	return ray.origin.Add(ray.direction.Mul(t))
}

//Intersect calculates the intersection between a ray and an object
func (ray *Ray) Intersect(object Object) []*Intersection {
	return object.Intersect(ray)
}
