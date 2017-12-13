package methods

import (
	"fmt"
	"math"
)

func hessian(p Point) [2][2]float64 {
	// Частные производные F второго порядка в точке p
	d2Fdx2 := 2.0
	d2Fdy2 := 2.0
	d2Fdxdy := 0.0

	return [2][2]float64{
		{d2Fdx2, d2Fdxdy},
		{d2Fdxdy, d2Fdy2},
	}
}

func det(m [2][2]float64) float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func invertedH(h [2][2]float64) [2][2]float64 {
	h[0][0], h[1][1] = h[1][1], h[0][0] // Адски убогое транспонирование

	d := det(h)

	if d == 0 {
		fmt.Errorf("Матрица гессе вырожденная: %v\n", h)
		return h
	}

	var inverted [2][2]float64
	for i, row := range h {
		for j := range row {
			inverted[i][j] = 1 / d * math.Pow(-1, float64(i+j+2)) * h[i][j]
		}
	}

	return inverted
}

// Критерий Сильвестра для матрицы 2х2
func isPositive(h [2][2]float64) bool {
	return h[0][0] > 0 && det(h) > 0
}

func multiply(p Point, m [2][2]float64) Point {
	return Point{
		m[0][0]*p.X + m[0][1]*p.Y,
		m[1][0]*p.X + m[1][1]*p.Y,
	}
}

func multiplyOnVal(m [2][2]float64, val float64) (res [2][2]float64) {
	for i, row := range m {
		for j, v := range row {
			res[i][j] = v * val
		}
	}

	return
}

func Newton(itCount int, curr Point, eps, lambda float64) Point {
	isAnswer := false

	for i := 0; i < itCount; i++ {
		grad := grad(curr)

		if vectorLength(grad) < eps {
			return curr
		}

		invH := invertedH(hessian(curr))

		var next Point
		if isPositive(invH) {
			d := multiply(grad, multiplyOnVal(invH, -1))
			next = curr.plus(d)
		} else {
			next = nextPoint(lambda, curr, grad) // Нахождение след. точки методом градиентного спуска.
			for F(next)-F(curr) >= 0 {           // Да-да, мне было лень его адапитровать
				lambda /= 2
				next = nextPoint(lambda, curr, grad)
			}
		}

		if vectorLength(next.deduct(curr)) < eps &&
			math.Abs(F(next)-F(curr)) < eps {
			if isAnswer {
				return next
			} else {
				isAnswer = true
			}
		} else if isAnswer {
			isAnswer = false
		}

		curr = next
	}

	return curr
}
