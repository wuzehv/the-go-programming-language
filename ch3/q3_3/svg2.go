// 对svg不了解，根据题干提到的高度，所以用z来判断了
package q3_2

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	cpicker       = 10 // 拾色器变化区间，太大会导致颜色变化过大
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Svg(writer io.Writer) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, c := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			var color string
			if c > 0 {
				// 这里做减法是因为数值越大颜色越应该接近纯红 ff0000
				color = red(cpicker - int(math.Floor(c*cpicker)))
			} else {
				color = blue(int(math.Floor(c*cpicker)) + cpicker)
			}

			fmt.Fprintf(writer, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style=\"fill: #%s;h:%g;\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color, c)
		}
	}
	fmt.Fprintf(writer, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.13f", math.Sin(r)/r), 64)
	return value
}

func red(i int) string {
	r := 0xFF0000
	return fmt.Sprintf("%x", r|i)
}

func blue(i int) string {
	b := 0x0000FF
	return fmt.Sprintf("%06x\n", b|i<<8)
}
