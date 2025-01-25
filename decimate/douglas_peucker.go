package decimate

import (
	"decimator/pkg/interfaces"
	"gonum.org/v1/gonum/mat"
)

func DouglasPeucker(geometry Geometry, points []Point) []mat.VecDense {
	var distance float64
	var distance_maximum float64
	distance_maximum = 0.0

	errors := ""
	// TODO add test to check if list of points have at least two points
	size_points := len(points)
	if size_points < 2 {
		errors += "Length of point list must be greater than one"
	}

	// TODO add test to check if all points have the same dimension
	for i, point := range points {
		if point.Dimension() != geometry.Dimension() {
			errors += fmt.Sprintf("All points must have the same dimension as the geometry. Point at position %v has dimension %v, but the geometry has dimension %v", i, point.Dimension(), geometry.Dimension())
		}
	}

	if errors != "" {
		log.Fatal(errors)
	}

	// TODO add test to check return when only to points are given
	if size_points == 2 {
		return points
	}

	dimension := points[0].Dimension()
	if dimension != 2 {
		geometry = interfaces.NewGeometry(dimension)
	}

	var index int
	for i := 1; i < size_points; i++ {
		distance = PerpendicularDistance(points[i], points[0], points[size_points-1])
		if distance > distance_maximum {
			index = i
			distance_maximum = distance
		}
	}
}
