package methods

import "math"

func diff(p1, p2 Point) float64 {
	return math.Abs(p1.X-p2.X) + math.Abs(p1.Y-p2.Y)
}

func PaulsMethod(curr Point, eps float64) Point {
	d1 := Point{1, 0}
	d2 := Point{0, 1}

	next := curr

	for {
		curr = next

		h1 := goldenSection(next, d1, eps)
		step := d1.multiply(h1)
		next = next.plus(step)

		h2 := goldenSection(next, d2, eps)
		step = d2.multiply(h2)
		next = next.plus(step)

		d1 = d1.multiply(h1)
		d2 = d2.multiply(h2)

		len := math.Abs(vectorLength(curr) - vectorLength(next))
		if len < eps {
			return next
		}
	}
}
