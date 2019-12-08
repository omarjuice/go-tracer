package main

//World creates an encapsulation of the objects and light
type World struct {
	lights  []*PointLight
	objects []Shape
}

//NewWorld ...
func NewWorld(lights []*PointLight, objects []Shape) *World {
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
	return NewWorld([]*PointLight{light}, []Shape{s1, s2})
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
func (world *World) ShadeHit(comps *Computation, remaining int) *Color {
	light := Black
	for i := 0; i < len(world.lights); i++ {
		light = light.Add(Lighting(
			comps.object.Material(),
			comps.object,
			world.lights[i],
			comps.overPoint,
			comps.eyev,
			comps.normalv,
			world.IsShadowed(comps.overPoint, i))).Add(world.ReflectedColor(comps, remaining))
	}
	return light
}

//ColorAt ...
func (world *World) ColorAt(ray *Ray, remaining int) *Color {
	hit := world.Intersect(ray).Hit()
	if hit == nil {
		return Black
	}
	comps := PrepareComputations(hit, ray)
	return world.ShadeHit(comps, remaining)
}

//ReflectedColor ...
func (world *World) ReflectedColor(comps *Computation, remaining int) *Color {
	if comps.object.Material().reflective == 0.0 || remaining < 1 {
		return Black
	}
	reflectRay := NewRay(comps.overPoint, comps.reflectv)
	color := world.ColorAt(reflectRay, remaining-1)

	return color.MulScalar(comps.object.Material().reflective)
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
	t                                         float64
	object                                    Shape
	point, eyev, normalv, reflectv, overPoint *Tuple
	inside                                    bool
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
	comps.reflectv = ray.direction.Reflect(comps.normalv)
	comps.overPoint = comps.point.Add(comps.normalv.Mul(EPSILON))
	return comps
}
