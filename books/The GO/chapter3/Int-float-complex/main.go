package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

// just hold positive value.
// unit8 represent an unsigned 8-but interger. meaning it can hold value ranging from 0-255.
// unit32 hold value from 0 to 4,294,967,295
// unit64 hold value ranging from 0 to 18,446,744,073,709,551,615

// int hold both negative and positive value.
// int8: This data type usually represents a signed 8-bit integer, meaning it can hold values ranging from -128 to 127. It's useful for conserving memory when you know your values will be within a small range.
// int32: This data type typically represents a signed 32-bit integer, which can hold values ranging from approximately -2.1 billion to 2.1 billion. It's commonly used for representing integers in a wide range of applications.
// int64: This data type generally represents a signed 64-bit integer, capable of holding values from approximately -9.2 quintillion to 9.2 quintillion. It's used when you need to work with very large integers or when dealing with values that may exceed the range of int32.

// float32 && float64
// float32 hold value from 1.4e-45 to 3.4e38
// float64 hold value from 4.9e-324 to 1.8e308

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/mandbrot", drawMandbrot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(genGraph()))
}

func drawMandbrot(w http.ResponseWriter, r *http.Request) {
	mainMandbrot(w, r)
}

func genGraph() string {
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if validateInf(ax) || validateInf(ay) || validateInf(bx) || validateInf(by) || validateInf(cx) || validateInf(cy) || validateInf(dx) || validateInf(dy) {
				continue
			}
			s = s + fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	s = s + fmt.Sprintf("</svg>")
	return s
}
func validateInf(x float64) bool {
	return math.IsNaN(x) && !math.IsInf(x, 0)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func mainMandbrot(out io.Writer, r *http.Request) {
	params := r.URL.Query()
	// const (
	// 	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	// 	width, height          = 1024, 1024
	// )
	xmin, _ := strconv.ParseFloat(params.Get("xmin"), 64)
	ymin, _ := strconv.ParseFloat(params.Get("ymin"), 64)
	xmax, _ := strconv.ParseFloat(params.Get("xmax"), 64)
	ymax, _ := strconv.ParseFloat(params.Get("ymax"), 64)
	width, _ := strconv.Atoi(params.Get("width"))
	height, _ := strconv.Atoi(params.Get("height"))

	if xmin == 0 {
		xmin = -2
	}
	if ymin == 0 {
		ymin = -2
	}
	if xmax == 0 {
		xmax = 2
	}
	if ymax == 0 {
		ymax = 2
	}
	if width == 0 {
		width = 1024
	}
	if height == 0 {
		height = 1024
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
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
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func lerp(v0, v1, t uint8) uint8 {
	return v0 + t*(v1-v0)
}
