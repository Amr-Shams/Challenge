package gmath 
import (
    
    "golang.org/x/exp/constraints"
)
type Number interface {
    constraints.Integer | constraints.Float
}

func Max[T Number](a, b T) T {
    if a > b {
        return a
    }
    return b
}
func Min[T Number](a, b T) T {
    return -Max(-a, -b)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
// find the LCM for an array of integers 
func LCMArray(integers []int) int {
    result := integers[0]
    for i := 1; i < len(integers); i++ {
        result = LCM(result, integers[i])
    }
    return result
}

