package methods

import "fmt"

func PatternSearch(start Point, eps, step float64) Point {
	lower, step := exploringSearch(start, step, eps)

	for step >= eps {
		distant := getDistantPoint(start, lower)

		s := step
		distant, s = exploringSearch(distant, s, eps)

		if F(distant) < F(lower) {
			if s < eps {
				return distant
			}

			start, lower = lower, distant
		} else {
			start = lower

			s2 := step
			lower, s2 = exploringSearch(start, s2, eps)

			if start.equal(lower) {
				return lower
			}
		}
	}

	fmt.Println("Unexpected behavior in pattern search method")
	return lower
}

func exploringSearch(point Point, step, eps float64) (Point, float64) {
	result := point

	for result == point {
		result = tryStep(point, step)

		if step < eps {
			return result, step
		}

		step /= 10
	}

	return result, step
}

func tryStep(point Point, step float64) Point {
	comparingVal := F(point)

	for i := 0; i < 2; i++ {
		for inverse := 0; inverse < 2; inverse++ {
			candidate := point

			offset := step
			if inverse != 0 {
				offset *= -1
			}

			if i == 0 {
				candidate.X += offset
			} else {
				candidate.Y += offset
			}

			candidateVal := F(candidate)
			if candidateVal < comparingVal {
				point = candidate
				comparingVal = candidateVal
			}
		}
	}

	return point
}

func getDistantPoint(p1, p2 Point) Point {
	return Point{distant(p1.X, p2.X), distant(p1.Y, p2.Y)}
}

func distant(a, b float64) float64 {
	return a + 2*(b-a)
}
