package methods

var e = [2][2]float64{
	{1, 0},
	{0, 1},
}

func Dfp(point Point, eps float64) Point {
	h := e
	g := grad(point)

	for {
		p := multiply(g, multiplyOnVal(h, -1))
		alpha := goldenSection(point, p, eps)

		next := p.multiply(alpha).plus(point)
		gNext := grad(next)

		if vectorLength(gNext) < eps {
			return next
		}

		v := next.deduct(point)
		u := gNext.deduct(g)

		a := calcA(v, u)
		b := calcB(h, u)

		h = sum(sum(h, a), b)
		point = next
		g = gNext
	}
}

func sum(a, b [2][2]float64) (res [2][2]float64) {
	for i, row := range a {
		for j, v := range row {
			res[i][j] = v + b[i][j]
		}
	}

	return
}

func columnOnRow(x, y Point) [2][2]float64 {
	return [2][2]float64{
		{y.X * x.X, y.X * x.Y},
		{y.Y * x.X, y.Y * x.Y},
	}
}

func calcA(v, u Point) [2][2]float64 {
	res := columnOnRow(v, v)

	uv := v.X*u.X + v.Y*u.Y

	return multiplyOnVal(res, 1/uv)
}

func calcB(h [2][2]float64, u Point) [2][2]float64 {
	hu := multiply(u, h)
	uth := Point{
		h[0][0]*u.X + h[0][1]*u.Y,
		h[1][0]*u.X + h[1][1]*u.Y,
	}

	mult := columnOnRow(hu, uth)
	dotProd := hu.X*u.X + hu.Y*u.Y

	return multiplyOnVal(mult, -1/dotProd)
}
