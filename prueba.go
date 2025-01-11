package main

import (
	"decimator/pkg/geom2d"
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// geometry := geom2d.NewEuclid()
	v1 := geom2d.NewVector(1, 0)
	v2 := geom2d.NewVector(0, 1)

	fmt.Println(mat.Formatted(v1))
	fmt.Println(mat.Formatted(v2))

}
