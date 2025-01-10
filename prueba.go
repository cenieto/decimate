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
	fmt.Println(v3.Components())
}
