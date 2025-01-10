package main

import (
	"decimator/pkg/geom2d"
	"errors"
	"fmt"
)

func main() {
	geometry := geom2d.NewGeometry()
	v1 := geom2d.Vector2D{1, 0}
	v2 := geom2d.Vector2D{0, 1}
	v3, error := geometry.CrossProduct(v1, v2)
	if error != nil {
		errors.New("Error in cross product calculation")
	}

	fmt.Printf("The cross product of %v over %v is %v\n",v1.Components(),v2.Components(),v3.Components())

	dimension, error := geometry.Norm(v2)

	fmt.Printf("The norm of v3 is %v\n", *dimension)

}}
