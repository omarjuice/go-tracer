package main

//Color is RGB color
type Color struct {
	r, g, b float64
}

//NewColor creates a new Color
func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

//Equals checks if a color is equal to another
func (c *Color) Equals(o *Color) bool {
	return floatEqual(c.r, o.r) && floatEqual(c.g, o.g) && floatEqual(c.b, o.b)
}

//Add adds a color with another
func (c *Color) Add(o *Color) *Color {
	return NewColor(
		c.r+o.r,
		c.g+o.g,
		c.b+o.b,
	)
}

//Sub subtracts two colors
func (c *Color) Sub(o *Color) *Color {
	return NewColor(
		c.r-o.r,
		c.g-o.g,
		c.b-o.b,
	)
}

//Mul blends two colors
func (c *Color) Mul(o *Color) *Color {
	return NewColor(
		c.r*o.r,
		c.g*o.g,
		c.b*o.b,
	)
}

//MulScalar multiplies a color by a scalar
func (c *Color) MulScalar(scalar float64) *Color {
	return NewColor(
		c.r*scalar,
		c.g*scalar,
		c.b*scalar,
	)
}

//Format turns the color into a string "red green blue" e.g. "255 128 16"
func (c *Color) Format() string {
	return floatToUint8String(c.r) + " " + floatToUint8String(c.g) + " " + floatToUint8String(c.b) + " "
}

//String formats a color as a string
func (c *Color) String() string {
	return "c(" + floatToString(c.r) + "," + floatToString(c.g) + "," + floatToString(c.b) + ")"
}
