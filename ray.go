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
func (ray *Ray) Intersect(object Shape) Intersections {
	intersections := object.Intersect(ray)

	return NewIntersections(intersections)
}

//Transform transforms a ray
func (ray *Ray) Transform(transformations ...Matrix) *Ray {
	return NewRay(
		ray.origin.Transform(transformations...),
		ray.direction.Transform(transformations...),
	)
}

//Equals checks ray equality
func (ray *Ray) Equals(other *Ray) bool {
	return ray.origin.Equals(other.origin) && ray.direction.Equals(other.direction)
}

//String formats Ray as a string
func (ray *Ray) String() string {
	return "r(" + ray.origin.String() + ray.direction.String() + ")"
}
