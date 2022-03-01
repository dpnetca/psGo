package main

import (
	"fmt"

	"github.com/dpnetca/psGo/functionDeepDive/simplemath"
)

type MathExpr = string

const (
	AddExpr = MathExpr("add")
	SubExpr = MathExpr("subtract")
)

func main() {
	fmt.Printf("Add: %f\n", simplemath.Add(2, 6))
	fmt.Printf("Sum: %f\n", simplemath.Sum(2, 6, 3, 4, 5))

	sv := simplemath.NewSemanticVersion(1, 2, 3)
	fmt.Println(sv.String())
	sv.IngrementPatch()
	fmt.Println(sv.String())
	sv.IngrementMinor()
	fmt.Println(sv.String())
	sv.IngrementMajor()
	fmt.Println(sv.String())

	func() {
		fmt.Println("anonymous function")
	}()

	af := func(name string) {
		fmt.Printf("anonymous function %s\n", name)
	}
	af("#2")
	af("#2")

	addExpr := mathExpression()
	fmt.Println(addExpr(1, 2))

	expr := anotherMathExpression(AddExpr)
	fmt.Println(expr(1, 2))
	expr = anotherMathExpression("subtract")
	fmt.Println(expr(1, 2))

}

func mathExpression() func(float64, float64) float64 {
	// return simplemath.Add
	return func(f1, f2 float64) float64 {
		return f1 + f2
	}
}

func anotherMathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubExpr:
		return simplemath.Subtract
	default:
		return func(f1, f2 float64) float64 {
			return 0
		}
	}
}
