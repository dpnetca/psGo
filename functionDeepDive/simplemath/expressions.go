package simplemath

// function names starting with capital are public, lower cased are private
func Add(p1, p2 float64) float64 {
	return p1 + p2
}
func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

// naming return variable in function definition you can use empty return statement
// but this is harder to read especially in longer functions
func Sum(values ...float64) (total float64) {
	total = 0.0
	for _, value := range values {
		total += value
	}
	return
}
