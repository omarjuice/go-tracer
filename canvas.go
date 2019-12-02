package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"
)

//Canvas is a canvas on which pixels are drawn
type Canvas struct {
	width, height    int
	pixels           [][]*Color
	originX, originY int
}

//NewCanvas creates a new canvas
func NewCanvas(width int, height int) *Canvas {
	canvas := &Canvas{width, height, make([][]*Color, height, height), 0, 0}
	for y := 0; y < height; y++ {
		canvas.pixels[y] = make([]*Color, width, width)
		for x := 0; x < width; x++ {
			canvas.pixels[y][x] = NewColor(0, 0, 0)
		}
	}
	return canvas
}

//SetOrigin sets the origin of the canvas
func (canvas *Canvas) SetOrigin(x, y int) {
	if !canvas.checkBounds(x, y) {
		return
	}

	canvas.originX = x
	canvas.originY = y
}

func (canvas *Canvas) checkBounds(x, y int) bool {
	r := true
	if y < 0 || y >= canvas.height {
		r = false
	}
	if x < 0 || x >= canvas.width {
		r = false
	}
	if !r {
		fmt.Println(x, y)
	}
	return r
}

//PixelAt returns the pixel at x and y
func (canvas *Canvas) PixelAt(x, y int) *Color {
	if !canvas.checkBounds(x, y) {
		return nil
	}
	return canvas.pixels[y][x]
}

//WritePixel writes a new pixel to the pixel
func (canvas *Canvas) WritePixel(x, y int, c *Color) {
	if !canvas.checkBounds(x, y) {
		return
	}
	canvas.pixels[y][x] = NewColor(c.r, c.g, c.b)
}

//Write writes a point to a canvas
func (canvas *Canvas) Write(point *Tuple, c *Color) {
	canvas.WritePixel(int(point.x)+canvas.originX, (int(point.y) + canvas.originY), c)
}

//ToPPM writes the canvas to a PPM file
func (canvas *Canvas) ToPPM(filename string) {
	lines := [][]rune{[]rune{}}

	for y := 0; y < canvas.height; y++ {
		for x := 0; x < canvas.width; x++ {
			color := canvas.pixels[y][x].Format()

			if len(color) > 69-len(lines[len(lines)-1]) {
				lines = append(lines, []rune{})
			}
			for _, ch := range color {
				lines[len(lines)-1] = append(lines[len(lines)-1], ch)
			}
		}
	}

	ppm := []string{"P3", strconv.Itoa(canvas.width) + " " + strconv.Itoa(canvas.height), "255"}

	for _, line := range lines {
		ppm = append(ppm, string(line))
	}
	ppm = append(ppm, "\n")

	filename += ".ppm"

	file, err := os.Open("/" + filename)

	if err != nil {
		file, _ = os.Create(filename)
	}
	defer file.Close()

	_, err = io.WriteString(file, strings.Join(ppm, "\n"))
	if err != nil {
		fmt.Println(err)
	}
}

func convert(colorVal float64) uint8 {
	r := uint8(max(min(colorVal*255, 255), 0))

	return r
}

//ToPNG converts a canvas to a PNG
func (canvas *Canvas) ToPNG(filename string) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{canvas.width, canvas.height}})

	for x := 0; x < canvas.width; x++ {
		for y := 0; y < canvas.height; y++ {
			c := canvas.PixelAt(x, y)
			color := color.RGBA{convert(c.r), convert(c.g), convert(c.b), 1}
			img.Set(x, y, color)
		}
	}

	filename += ".png"

	file, err := os.Open("/" + filename)

	if err != nil {
		file, _ = os.Create(filename)
	}
	defer file.Close()

	png.Encode(file, img)

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
