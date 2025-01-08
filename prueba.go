package main

import (
	"decimator/pkg/geom3d"
	"errors"
	"fmt"
)

func main() {
	v1 := geom3d.Vector3D{1, 0, 0}
	v2 := geom3d.Vector3D{0, 1, 0}
	v3, error := v1.CrossProduct(v2)
	if error != nil {
		errors.New("Error in cross product calculation")
	}
	fmt.Println(v3.Components())
}
