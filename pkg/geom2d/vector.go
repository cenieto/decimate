package geom2d

import (
	"decimator/pkg/geom3d"
	"decimator/pkg/interfaces"
	"errors"
)

type Vector2D struct {
	X, Y float64
}

func NewVector(x, y float64) Vector2D {
	return Vector2D{
		X: x,
		Y: y,
	}
}

func (v Vector2D) Dimension() int {
	return 2
}

func (v Vector2D) Components() []float64 {
	return []float64{v.X, v.Y}
}

func (v Vector2D) CrossProduct(other interfaces.Vectorial) (interfaces.Vectorial, error) {
	if other.Dimension() != 2 {
		return nil, errors.New("Cross product is not defined for vectors of different dimensions")
	}

	v1 := v.Components()
	v2 := other.Components()
	result := geom3d.Vector3D{
		X: 0.0,
		Y: 0.0,
		Z: v1[0]*v2[1] - v1[1]*v2[0],
	}

	return result, nil
}
