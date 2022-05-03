package golang_generics

import (
	"fmt"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Rizki")
	fmt.Println(result)

	var resultNumber int = Length[int](26)
	fmt.Println(resultNumber)

	var resultBoolean bool = Length[bool](false)
	fmt.Println(resultBoolean)
}
