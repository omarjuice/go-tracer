package main

import (
	"math"
)

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
func (world *World) Intersect(ray *Ray) Intersections {
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
	material := comps.object.Material()
	reflectance := 1.0
	refractance := 1.0
	if material.reflective > 0 && material.transparency > 0 {
		reflectance = comps.Schlick()
		refractance = 1 - reflectance
	}
	for i := 0; i < len(world.lights); i++ {

		light = light.Add(
			Lighting(
				comps.object.Material(),
				comps.object,
				world.lights[i],
				comps.overPoint,
				comps.eyev,
				comps.normalv,
				world.IsShadowed(comps.overPoint, i)),
		).Add(
			world.ReflectedColor(comps, remaining).MulScalar(reflectance),
		).Add(
			world.RefractedColor(comps, remaining).MulScalar(refractance),
		)
	}
	return light
}

//ColorAt ...
func (world *World) ColorAt(ray *Ray, remaining int) *Color {
	xs := world.Intersect(ray)
	hit := xs.Hit()
	if hit == nil {
		return Black
	}
	comps := PrepareComputations(hit, ray, xs)
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

//RefractedColor ...
func (world *World) RefractedColor(comps *Computation, remaining int) *Color {
	if comps.object.Material().transparency == 0.0 || remaining < 1 {
		return Black
	}
	nRatio := comps.n1 / comps.n2

	cosI := comps.eyev.Dot(comps.normalv)

	sin2t := square(nRatio) * (1 - square(cosI))

	if sin2t > 1.0 {
		return Black
	}

	cosT := math.Sqrt(1.0 - sin2t)

	direction := comps.normalv.Mul(nRatio*cosI - cosT).Sub(comps.eyev.Mul(nRatio))

	refractRay := NewRay(comps.underPoint, direction)

	color := world.ColorAt(refractRay, remaining-1).MulScalar(comps.object.Material().transparency)

	return color
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
	t, n1, n2                                             float64
	object                                                Shape
	point, eyev, normalv, reflectv, overPoint, underPoint *Tuple
	inside                                                bool
}

//PrepareComputations ...
func PrepareComputations(hit *Intersection, ray *Ray, xs Intersections) *Computation {
	point := ray.Position(hit.t)
	comps := &Computation{
		t:       hit.t,
		object:  hit.object,
		point:   point,
		eyev:    ray.direction.Negate(),
		normalv: hit.object.NormalAt(point),
		inside:  false,
	}
	if comps.normalv.Dot(comps.eyev) < 0 {
		comps.inside = true
		comps.normalv = comps.normalv.Negate()
	}

	comps.reflectv = ray.direction.Reflect(comps.normalv)
	comps.overPoint = comps.point.Add(comps.normalv.Mul(EPSILON))
	comps.underPoint = comps.point.Sub(comps.normalv.Mul(EPSILON))

	containers := []Shape{}

	for _, inters := range xs {
		if inters == hit {
			if len(containers) == 0 {
				comps.n1 = 1
			} else {
				comps.n1 = containers[len(containers)-1].Material().refractiveIndex
			}
		}
		if !removeIfContains(&containers, inters.object) {
			containers = append(containers, inters.object)
		}
		if inters == hit {
			if len(containers) == 0 {
				comps.n2 = 1
			} else {
				comps.n2 = containers[len(containers)-1].Material().refractiveIndex
			}
			break
		}
	}

	return comps
}

//Schlick returns the reflectance, represents what fraction of light is reflected given surface info and the Intersections.Hit()
func (comps *Computation) Schlick() float64 {
	cos := comps.eyev.Dot(comps.normalv)
	if comps.n1 > comps.n2 {
		n := comps.n1 / comps.n2
		sin2T := square(n) * (1.0 - square(cos))
		if sin2T > 1.0 {
			return 1.0
		}
		cosT := math.Sqrt(1.0 - sin2T)
		cos = cosT
	}
	ro := square((comps.n1 - comps.n2) / (comps.n1 + comps.n2))

	return ro + (1-ro)*math.Pow(1-cos, 5)
}

func removeIfContains(containers *[]Shape, obj Shape) bool {
	C := *containers
	for i, shape := range C {
		if shape == obj {
			for i < len(C)-1 {
				C[i] = C[i+1]
				i++
			}
			*containers = C[:len(C)-1]
			return true
		}
	}
	return false
}
