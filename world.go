package main

//World creates an encapsulation of the objects and light
type World struct {
	light   *PointLight
	objects []Object
}

//NewWorld ...
func NewWorld(light *PointLight, objects []Object) *World {
	return &World{light, objects}
}

//DefaultWorld ...
func DefaultWorld() *World {
	light := NewPointLight(Point(-10, 10, -10), NewColor(1, 1, 1))
	s1 := NewSphere()
	s1.material.color = NewColor(0.8, 1.0, .6)
	s1.material.diffuse = 0.7
	s1.material.specular = 0.2

	s2 := NewSphere()
	s2.SetTransform(Scaling(.5, .5, .5))
	return NewWorld(light, []Object{s1, s2})
}

//Intersect a ray with the world
func (world *World) Intersect(ray *Ray) *Intersections {
	intersections := []*Intersection{}

	for _, object := range world.objects {
		intersections = append(intersections, object.Intersect(ray)...)
	}

	xs := NewIntersections(intersections)

	return xs

}

//Computation ...
type Computation struct {
	t                    float64
	object               Object
	point, eyev, normalv *Tuple
	inside               bool
}

//PrepareComputations ...
func PrepareComputations(intersection *Intersection, ray *Ray) *Computation {
	point := ray.Position(intersection.t)
	comps := &Computation{
		t:       intersection.t,
		object:  intersection.object,
		point:   point,
		eyev:    ray.direction.Negate(),
		normalv: intersection.object.NormalAt(point),
		inside:  false,
	}
	if comps.normalv.Dot(comps.eyev) < 0 {
		comps.inside = true
		comps.normalv = comps.normalv.Negate()
	}
	return comps
}
