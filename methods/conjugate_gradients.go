package methods

import "math"

func ConjugateGradients(n int, curr Point, eps float64) Point {
	gradCurr := grad(curr)
	p := gradCurr.multiply(-1)

	for i := 0; ; i++ {
		alpha := goldenSection(curr, p, eps)
		next := curr.plus(p.multiply(alpha))

		gradNext := grad(next)
		if vectorLength(gradNext) < eps {
			return next
		} else if i+1 == n {
			curr = next
			gradCurr = gradNext
			p = gradCurr.multiply(-1)

			continue
		}

		b := math.Pow(vectorLength(gradNext), 2) / math.Pow(vectorLength(gradCurr), 2)
		p = gradNext.
			multiply(-1).
			plus(p.multiply(b))

		curr = next
	}
}
