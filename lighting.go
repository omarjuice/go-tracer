package main

import "math"

//PointLight is a light source with no size, existing at a single point in space
type PointLight struct {
	position  *Tuple
	intensity *Color
}

//NewPointLight returns a PointLight
func NewPointLight(position *Tuple, intensity *Color) *PointLight {
	return &PointLight{position, intensity}
}

//Material encapsulates the given attributes of the Phong reflection model
type Material struct {
	color *Color

	ambient, diffuse, specular, shininess float64
}

//DefaultMaterial ...
func DefaultMaterial() *Material {
	return NewMaterial(White, 0.1, .9, .9, 200.0)
}

//NewMaterial creates a new Materials
func NewMaterial(color *Color, ambient, diffuse, specular, shininess float64) *Material {
	return &Material{color, ambient, diffuse, specular, shininess}
}

//Lighting computes lighting
func Lighting(material *Material, light *PointLight, point, eyev, normalv *Tuple) *Color {
	effectiveColor := material.color.Mul(light.intensity)

	lightv := light.position.Sub(point).Normalize()

	ambient := effectiveColor.MulScalar(material.ambient)

	lightDotNormal := lightv.Dot(normalv)

	diffuse := Black
	specular := Black

	if lightDotNormal >= 0 {
		diffuse = effectiveColor.MulScalar(material.diffuse).MulScalar(lightDotNormal)

		reflectv := lightv.Negate().Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, material.shininess)
			specular = light.intensity.MulScalar(material.specular).MulScalar(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
