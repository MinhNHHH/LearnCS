package complexnumber

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func draw(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colorsize(n, contrast)
		}
	}
	return color.Black
}

func colorsize(n uint8, contrast int) color.Color {
	//
	blue := contrast * 255
	red := 255 - blue
	green := lerp(n, uint8(red), uint8(blue))
	return color.RGBA{uint8(red), uint8(green), uint8(blue), 255}
}

func lerp(percentage uint8, hue0 uint8, hue1 uint8) uint8 {
	return (percentage*(hue1-hue0) + hue0)
}
