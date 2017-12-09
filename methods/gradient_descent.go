package methods

import "math"

// Метод градиентного спуска

func GradientDescent(itCount int, curr Point, eps, lambda float64) Point {
	isAnswer := false
	for i := 0; i < itCount; i++ {
		grad := grad(curr)

		if vectorLength(grad) < eps {
			return curr
		}

		next := nextPoint(lambda, curr, grad)
		for F(next)-F(curr) >= 0 {
			lambda /= 2
			next = nextPoint(lambda, curr, grad)
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

func nextPoint(lambda float64, point, grad Point) Point {
	return Point{
		point.X - lambda*grad.X,
		point.Y - lambda*grad.Y,
	}
}
