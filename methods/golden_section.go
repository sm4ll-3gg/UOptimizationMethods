package methods

import "math"

const (
	BEGIN = 0.0
	END   = 1000.0
)

var PHI = (3 - math.Sqrt(5)) / 2

func goldenSection(point, p Point, eps float64) float64 {
	left := BEGIN
	right := END

	al1 := left + PHI*(right-left)
	al2 := left + right - al1

	for {
		f1 := F(point.plus(p.multiply(al1)))
		f2 := F(point.plus(p.multiply(al2)))

		if f1 <= f2 {
			right = al2
			al2 = al1
			al1 = left + right - al1
		} else {
			left = al1
			al1 = al2
			al2 = left + right - al2
		}

		if math.Abs(left-right) < eps {
			return (left + right) / 2
		}
	}
}
