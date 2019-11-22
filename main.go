package main

import "fmt"

func main() {
	// fmt.Println("a"[-1])
	canvas := NewCanvas(900, 550)
	start := Point(0, 1, 0)
	velocity := Vector(1, 1.8, 0).Normalize().Mul(11.25)
	p := &Projectile{start, velocity}

	gravity := Vector(0, -0.1, 0)
	wind := Vector(-0.01, 0, 0)

	e := &Environment{gravity, wind}

	color := NewColor(0, 1, 0)

	p.WriteToCanvas(canvas, color)
	i := 0
	for p.position.y >= 0.0 {
		p = Tick(e, p)
		p.WriteToCanvas(canvas, color)
		i++
	}
	fmt.Println(i)

	canvas.ToPPM("render")
}
