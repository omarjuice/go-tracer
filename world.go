package main

//World creates an encapsulation of the objects and light
type World struct {
	lights  []*PointLight
	objects []Object
}

//NewWorld ...
func NewWorld(lights []*PointLight, objects []Object) *World {
	return &World{lights, objects}
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
	return NewWorld([]*PointLight{light}, []Object{s1, s2})
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

//ShadeHit returns the color encapsulated by comps in the given world
func (world *World) ShadeHit(comps *Computation) *Color {
	light := Lighting(comps.object.Material(), world.lights[0], comps.point, comps.eyev, comps.normalv, world.IsShadowed(comps.overPoint, 0))
	for i := 1; i < len(world.lights); i++ {
		light = light.Add(Lighting(comps.object.Material(), world.lights[i], comps.point, comps.eyev, comps.normalv, world.IsShadowed(comps.overPoint, i)))
	}
	return light
}

//ColorAt ...
func (world *World) ColorAt(ray *Ray) *Color {
	hit := world.Intersect(ray).Hit()
	if hit == nil {
		return Black
	}
	comps := PrepareComputations(hit, ray)
	return world.ShadeHit(comps)
}

//IsShadowed returns whether a point is in a shadow
func (world *World) IsShadowed(point *Tuple, light int) bool {
	v := world.lights[light].position.Sub(point)
	distance := v.Magnitude()
	direction := v.Normalize()

	ray := NewRay(point, direction)

	intersections := world.Intersect(ray)

	hit := intersections.Hit()

	return hit != nil && hit.t < distance
}

//Computation ...
type Computation struct {
	t                               float64
	object                          Object
	point, eyev, normalv, overPoint *Tuple
	inside                          bool
}

//PrepareComputations ...
func PrepareComputations(intersection *Intersection, ray *Ray) *Computation {
	point := ray.Position(intersection.t)
	comps := &Computation{
		t:         intersection.t,
		object:    intersection.object,
		point:     point,
		overPoint: nil,
		eyev:      ray.direction.Negate(),
		normalv:   intersection.object.NormalAt(point),
		inside:    false,
	}
	if comps.normalv.Dot(comps.eyev) < 0 {
		comps.inside = true
		comps.normalv = comps.normalv.Negate()
	}
	comps.overPoint = comps.point.Add(comps.normalv.Mul(EPSILON))
	return comps
}
