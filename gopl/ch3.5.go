package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

func dispRunPercent(count, total int) {
	fmt.Printf("\rRun %d%%", count*100/total)
}

func drawMandelbrot() {
	const (
		xMin, yMin, xMax, yMax = -2, -2, 2, 2
		border                 = 2000
		width, height          = border, border
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dispCount := 0
	dispTotal := height * width
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
			dispCount++
			dispRunPercent(dispCount, dispTotal)
		}
	}
	pngFile, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err.Error())
	}
	defer pngFile.Close()
	_ = png.Encode(pngFile, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 100
		contrast   = 15
	)

	rand.Seed(time.Now().Unix())
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				A: 255,
				B: 255 - contrast*n,
				G: 127 - contrast*n,
				R: 64 + contrast*n,
			}
		}
	}
	return color.Black
}

func main() {
	drawMandelbrot()
}
