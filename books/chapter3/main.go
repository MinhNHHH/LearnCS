package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
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
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(genGraph()))
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

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func color(z)
