package decimate

import (
	"decimator/pkg/geom2d"
	"fmt"
	"testing"
)

func TestDouglasPeuckerShortInput(t *testing.T) {
	geometry := geom2d.NewEuclid()
	var points []*geom2d.Point2D
	points[0] = geometry.NewPoint(0.0, 0.0)

	result, errors := DouglasPeucker(geometry, points)
	if errors == nil {
		t.Errorf("DouglasPeucker request len(%) > 1. Input check is incorrectly done.")
	}

	fmt.Println(result)
}
