package methods

import "math"

// Метод градиентного спуска

type Point struct {
	X float64
	Y float64
}

func (p Point) deduct(p1 Point) Point {
	return Point{p.X - p1.X, p.Y - p1.Y}
}

func GradientDescent(itCount int, curr Point, eps, lambda float64) Point {
	for i := 0; i < itCount; i++ {
		grad := Grad(curr) // start

		if vectorLength(grad) < eps {
			return curr
		}

		next := nextPoint(lambda, curr, grad);
		for step := lambda / 2; F(next) - F(curr) >= 0; step /= 2 {
			next = nextPoint(step, curr, grad)
		}

		if vectorLength(next.deduct(curr)) < eps &&
			math.Abs(F(next) - F(curr)) < eps {
			return next
		}

		curr = next
	}

	return curr
}

func nextPoint(lambda float64, point, grad Point) Point {
	return Point{
		point.X - lambda*grad.X,
		point.Y - lambda*grad.Y,
	}
}
