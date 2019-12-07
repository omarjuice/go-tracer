package main

import "math"

type getColorFunc func([]*Color, *Tuple) *Color

//Pattern ...
type Pattern struct {
	colors    [][]*Color
	funcs     []getColorFunc
	transform Matrix
}

var stripeFunc = func(colors []*Color, p *Tuple) *Color {
	return colors[(int(abs(p.x)))%len(colors)]
}

var gradientFunc = func(colors []*Color, p *Tuple) *Color {
	dist := colors[1].Sub(colors[0])
	frac := p.x - math.Floor(p.x)

	return colors[0].Add(dist.MulScalar(frac))
}

var ringFunc = func(colors []*Color, p *Tuple) *Color {
	return colors[int(math.Floor(math.Sqrt(square(p.x)+square(p.z))))%len(colors)]
}

var checkersFunc = func(colors []*Color, p *Tuple) *Color {
	if (int(p.x)+int(p.y)+int(p.z))%2 == 0 {
		return colors[0]
	}
	return colors[1]

}

//NewPattern ...
func NewPattern(colors [][]*Color, getColor ...getColorFunc) *Pattern {
	return &Pattern{colors, getColor, NewIdentityMatrix()}
}

//ColorAt ...
func (pattern *Pattern) ColorAt(p *Tuple) *Color {
	color := Black
	for i := 0; i < len(pattern.funcs); i++ {
		color = color.Add(pattern.funcs[i](pattern.colors[i], p))
	}
	return color
}

//ColorAtObject ...
func (pattern *Pattern) ColorAtObject(object Shape, point *Tuple) *Color {
	op := object.Transform().MulTuple(point)
	pp := pattern.transform.MulTuple(op)

	return pattern.ColorAt(pp)
}

//SetTransform ...
func (pattern *Pattern) SetTransform(transform Matrix) {
	pattern.transform = transform.Inverse()
}

//StripePattern ...
func StripePattern(colors ...*Color) *Pattern {

	return NewPattern([][]*Color{colors}, stripeFunc)
}

//GradientPattern ...
func GradientPattern(a, b *Color) *Pattern {
	return NewPattern([][]*Color{[]*Color{a, b}}, gradientFunc)
}

//RingPattern ...
func RingPattern(colors ...*Color) *Pattern {
	return NewPattern([][]*Color{colors}, ringFunc)
}

//CheckersPattern ...
func CheckersPattern(a, b *Color) *Pattern {
	return NewPattern([][]*Color{[]*Color{a, b}}, checkersFunc)
}

//PatternChain chains patterns together
func PatternChain(patterns ...*Pattern) *Pattern {
	colors := [][]*Color{}
	funcs := []getColorFunc{}
	transform := NewIdentityMatrix()
	for _, p := range patterns {
		for _, cs := range p.colors {
			colors = append(colors, cs)
		}
		for _, f := range p.funcs {
			funcs = append(funcs, f)
		}
		transform = transform.MulMatrix(p.transform)
	}
	pat := NewPattern(colors, funcs...)
	pat.SetTransform(transform)
	return pat
}
