package main

import "fmt"
import "./methods"

var (
	x0 = 0.0
	y0 = 0.0

	eps = 0.01

	lambda = 1.0
)

//func init() {
//	fmt.Println("Введите x0: ")
//	fmt.Scan(&x0)
//
//	fmt.Println("Введите y0: ")
//	fmt.Scan(&y0)
//
//	fmt.Println("Введите точность: ")
//	fmt.Scan(&eps)
//}

func main() {
	start := methods.Point{x0, y0}
	p := methods.GradientDescent(10, start, eps, lambda)
	fmt.Printf("x = %f; y = %f;\nf(x,y) = %f", p.X, p.Y, methods.F(p))
}
