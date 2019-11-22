package main

//Projectile ...
type Projectile struct {
	position *Tuple
	velocity *Tuple
}

//Environment ...
type Environment struct {
	gravity *Tuple
	wind    *Tuple
}

//Tick ...
func Tick(env *Environment, p *Projectile) *Projectile {
	position := p.position.Add(p.velocity)
	velocity := p.velocity.Add(env.gravity).Add(env.wind)

	return &Projectile{position, velocity}
}

//WriteToCanvas ...
func (p *Projectile) WriteToCanvas(canvas *Canvas, color *Color) {
	canvas.WritePixel(int(p.position.x), canvas.height-int(p.position.y), color)
}
