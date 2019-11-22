package main

import "testing"

func TestNewCanvas(t *testing.T) {
	h := 10
	w := 20
	canvas := NewCanvas(w, h)

	if len(canvas.pixels) != h {
		t.Errorf("NewCanvas: height of canvas should be %v but got %v", h, len(canvas.pixels))
	}
	defaultColor := NewColor(0, 0, 0)

	for y, row := range canvas.pixels {
		if len(row) != w {
			t.Errorf("NewCanvas: width of canvas should be %v but got %v", w, len(row))
		}
		for x, px := range row {
			if !px.Equals(defaultColor) {
				t.Errorf("NewCanvas: pixel at %v,%v is not of default color", x, y)
			}
		}
	}
}

func TestCanvasPixels(t *testing.T) {
	h := 10
	w := 20
	canvas := NewCanvas(h, w)

	red := NewColor(1, 0, 0)
	x, y := 2, 3
	canvas.WritePixel(x, y, red)
	px := canvas.PixelAt(x, y)
	if !px.Equals(red) {
		t.Errorf("CanvasPixels: pixel %v at %v,%v should be %v", px.String(), x, y, red.String())
	}
}
