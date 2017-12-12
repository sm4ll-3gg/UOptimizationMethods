package methods

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) deduct(p1 Point) Point {
	return Point{p.X - p1.X, p.Y - p1.Y}
}

func (p Point) multiply(x float64) Point {
	return Point{x * p.X, x * p.Y}
}

func (p Point) plus(p1 Point) Point {
	return Point{p.X + p1.X, p.Y + p1.Y}
}

func (p Point) equal(other Point) bool {
	return (p.X == other.X) && (p.Y == other.Y)
}

// Функция для которой находится минимум
func F(p Point) float64 {
	return math.Pow(p.X, 3) + math.Pow(p.Y, 2) - 3*p.X - 2*p.Y + 2
}

// Частная производная функции F по X в точке p
func dFdx(p Point) float64 {
	return 3*math.Pow(p.X, 2) - 3
}

// Частная производная функции F по Y в точке p
func dFdy(p Point) float64 {
	return 2*p.Y - 2
}

// Градиент функции F в точке p
func grad(p Point) Point {
	return Point{
		dFdx(p),
		dFdy(p),
	}
}

// Модуль градиента в точке
func vectorLength(grad Point) float64 {
	return math.Sqrt(math.Pow(grad.X, 2) + math.Pow(grad.Y, 2))
}
