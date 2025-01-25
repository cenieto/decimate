package decimate

import (
	"decimator/pkg/interfaces"
	"errors"
)

func DouglasPeucker(geometry interfaces.Geometry, points []interfaces.Point) ([]interfaces.Point, error) {
	// var distance float64
	// var distance_maximum float64
	// distance_maximum = 0.0

	errorMsgs := ""
	// TODO add test to check if list of points have at least two points
	size_points := len(points)
	if size_points < 2 {
		errorMsgs += "Length of point list must be greater than one"
	}

	if len(errorMsgs) > 0 {
		return nil, errors.New(errorMsgs)
	}

	return nil, nil
}
