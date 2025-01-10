package geom2d

import (
	"decimator/pkg/interfaces"
	"decimator/pkg/geom3d"
	"errors"
	"fmt"
	"math"
)

type Geometry2D struct {
}

func NewGeometry() Geometry2D {
	return Geometry2D{}
}

func (g Geometry2D) CrossProduct(args ...interfaces.Vectorial) (interfaces.Vectorial, error) {
	errorMsg := ""
	for _, arg := range args {
		if arg.Dimension() != 2 {
			errorMsg += fmt.Sprintf("Geometry2D is only defined for vectors of dimension 2, %v is of dimension %v\n", arg, arg.Dimension())
		}
	}

	if len(errorMsg) > 1 {
		return nil, errors.New(errorMsg)
	}

	v1 := args[0].Components()
	v2 := args[1].Components()
	
	result := geom3d.Vector3D{
		X: 0.0,
		Y: 0.0,
		Z: v1[0]*v2[1] - v1[1]*v2[0],
	}

	return result, nil
}

func (g Geometry2D) Norm(v interfaces.Vectorial) (*float64, error){
	dimension := v.Dimension()
	if dimension != 2{
		errorMsg := fmt.Sprintf("Geometry2D is only defined for vectors of dimension 2, v is of dimension %v", dimension)
		return nil, errors.New(errorMsg)
	}

	v1 := v.Components()
	result := math.Sqrt(v1[0]*v1[0]+ v1[1]*v1[1])
	return &result, nil
}
