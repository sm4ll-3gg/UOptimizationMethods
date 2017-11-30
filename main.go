package main

import "fmt"
import "./methods"

const (
	x0 = 0.0
	y0 = 0.0

	eps = 0.1

	lambda = 1.0
)

func main() {
	start := methods.Point{x0, y0}
	p := methods.GradientDescent(4, start, eps, lambda)
	fmt.Printf("x = %f; y = %f;\nf(x,y) = %f", p.X, p.Y, methods.F(p))
}
