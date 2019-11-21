package main

import "strings"

//Canvas is a canvas on which pixels are drawn
type Canvas struct {
	height, width int
	pixels        [][]*Color
}

//NewCanvas creates a new canvas
func NewCanvas(height int, width int) *Canvas {
	canvas := &Canvas{height, width, make([][]*Color, height, height)}
	for y := 0; y < height; y++ {
		canvas.pixels[y] = make([]*Color, width, width)
		for x := 0; x < width; x++ {
			canvas.pixels[y][x] = NewColor(0, 0, 0)
		}
	}
	return canvas
}

//PixelAt returns the pixel at x and y
func (canvas *Canvas) PixelAt(x, y int) *Color {
	return canvas.pixels[y][x]
}

//WritePixel writes a new pixel to the pixel
func (canvas *Canvas) WritePixel(x, y int, c *Color) {
	canvas.pixels[y][x] = NewColor(c.r, c.g, c.b)
}

//String converts a canvas to a string
func (canvas *Canvas) String() string {
	pixels := make([]string, canvas.height, canvas.height)

	for y := 0; y < canvas.height; y++ {
		row := make([]string, canvas.width, canvas.width)
		for x := 0; x < canvas.width; x++ {
			row[x] = canvas.PixelAt(x, y).String()
		}
		pixels[y] = strings.Join(row, ",")
	}

	return strings.Join(pixels, "\n")
}
