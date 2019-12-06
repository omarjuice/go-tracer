package main

import (
	"math"
)

//Shape is any object in a scene
type Shape interface {
	SetMaterial(*Material)
	SetTransform(Matrix)
	Transform() Matrix
	Material() *Material
	Intersect(*Ray) []*Intersection
	NormalAt(*Tuple) *Tuple
}

//Sphere object...
type Sphere struct {
	origin    *Tuple
	transform Matrix
	material  *Material
}

//NewSphere creates a new sphere
func NewSphere() *Sphere {
	return &Sphere{Point(0, 0, 0), IdentityMatrix, DefaultMaterial()}
}

//Material returns the material of a Sphere
func (sphere *Sphere) Material() *Material {
	return sphere.material
}

//Transform returns the transformation
func (sphere *Sphere) Transform() Matrix {
	return sphere.transform
}

//SetTransform sets the spheres transformation
func (sphere *Sphere) SetTransform(transformation Matrix) {
	sphere.transform = transformation.Inverse()
}

//SetMaterial sets the spheres material
func (sphere *Sphere) SetMaterial(material *Material) {
	sphere.material = material
}

//Intersect computes the intersection between a sphere and a ray
func (sphere *Sphere) Intersect(ray *Ray) []*Intersection {
	ray = ray.Transform(sphere.transform)
	sphereToRay := ray.origin.Sub(sphere.origin)
	a := ray.direction.Dot(ray.direction)
	b := 2 * ray.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return []*Intersection{}
	}
	sqrtDisc := math.Sqrt(discriminant)
	div := (2 * a)
	t1 := (-b - sqrtDisc) / div
	t2 := (-b + sqrtDisc) / div
	return []*Intersection{NewIntersection(t1, sphere), NewIntersection(t2, sphere)}

}

//NormalAt calculates the normal(vector perpendicular to the surface) at a given point
func (sphere *Sphere) NormalAt(point *Tuple) *Tuple {
	objectPoint := sphere.transform.MulTuple(point)
	objectNormal := objectPoint.Sub(sphere.origin)
	worldNormal := sphere.transform.Transpose().MulTuple(objectNormal)

	worldNormal.w = 0.0
	return worldNormal.Normalize()
}

//Plane Shape
type Plane struct {
	transform Matrix
	material  *Material
}

//NewPlane ...
func NewPlane() *Plane {
	return &Plane{NewIdentityMatrix(), DefaultMaterial()}

}

//Transform ...
func (plane *Plane) Transform() Matrix {
	return plane.transform
}

//Material ...
func (plane *Plane) Material() *Material {
	return plane.material
}

//SetTransform ...
func (plane *Plane) SetTransform(transform Matrix) {
	plane.transform = transform.Inverse()
}

//SetMaterial ...
func (plane *Plane) SetMaterial(material *Material) {
	plane.material = material
}

//Intersect ...
func (plane *Plane) Intersect(ray *Ray) []*Intersection {
	if abs(ray.direction.y) < EPSILON {
		return []*Intersection{}
	}
	ray = ray.Transform(plane.transform)
	t := -ray.origin.y / ray.direction.y

	return []*Intersection{NewIntersection(t, plane)}
}

//NormalAt ...
func (plane *Plane) NormalAt(point *Tuple) *Tuple {
	return Vector(0, 1, 0)
}
