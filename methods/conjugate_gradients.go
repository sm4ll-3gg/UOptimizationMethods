package methods

import "math"

func ConjugateGradients(n int, curr Point, eps float64) Point {
	gradCurr := grad(curr)
	p := gradCurr.multiply(-1)

	i := 0
	for {
		alpha := goldenSection(curr, p, eps)

		next := p.multiply(alpha).plus(curr)
		gradNext := grad(next)

		if vectorLength(gradNext) < eps {
			return next
		} else if i+1 == n {
			curr = next
			gradCurr = gradNext

			p = gradCurr.multiply(-1)
			i = 0

			continue
		}

		b := math.Pow(vectorLength(gradNext), 2) / math.Pow(vectorLength(gradCurr), 2)
		p = gradNext.multiply(-1).plus(p.multiply(b))

		i++
	}
	//gradCurr := grad(curr)
	//p := gradCurr.multiply(-1)
	//
	//i := 0
	//for {
	//	alpha := goldenSection(curr, p, eps)
	//	next := p.multiply(alpha).plus(curr)
	//
	//	gradNext := grad(next)
	//	if vectorLength(gradNext) < eps {
	//		return next
	//	} else if i+1 == n {
	//		curr = next
	//		gradCurr = gradNext
	//
	//		// reset
	//		p = gradCurr.multiply(-1)
	//		i = 0
	//
	//		continue
	//	}
	//
	//	b := math.Pow(vectorLength(gradNext), 2) / math.Pow(vectorLength(gradCurr), 2)
	//	//p = gradNext.
	//	//	multiply(-1).
	//	//	plus(p.multiply(b))
	//	p = gradCurr.multiply(b).plus(gradNext)
	//
	//	curr = next
	//	i++
	//}
}
