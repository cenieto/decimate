package geom3d

import (
	"decimator/pkg/interfaces"
	"errors"
)

type Vector3D struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector3D {
	return Vector3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v Vector3D) Dimension() int {
	return 3
}

func (v Vector3D) Components() []float64 {
	return []float64{v.X, v.Y, v.Z}
}

func (v Vector3D) CrossProduct(other interfaces.Vectorial) (interfaces.Vectorial, error) {
	if other.Dimension() != 3 {
		return nil, errors.New("Cross product is not defined for vectors of different dimensions")
	}

	v1 := v.Components()
	v2 := other.Components()
	result := Vector3D{
		X: v1[1]*v2[2] - v1[2]*v2[1],
		Y: -v1[0]*v2[2] + v1[2]*v2[0],
		Z: v1[0]*v2[1] - v1[1]*v2[0],
	}

	return result, nil
}
