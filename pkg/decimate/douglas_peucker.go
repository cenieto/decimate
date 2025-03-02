// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package decimate

import (
	"errors"
	"fmt"
	"github.com/cenieto/decimate/pkg/interfaces"
	"github.com/cenieto/decimate/pkg/primitives"
)

// Decimate is a struct that represents a decimate operation.
type Decimate struct {
	Geometry interfaces.Geometry
}

// NewDecimate creates and returns a new instance of Decimate.
//
// Parameters:
//   - geometry (interfaces.Geometry): The geometry to be decimated.
//
// Returns:
//   - *Decimate: A new instance of the decimate operation.
func NewDecimate(geometry interfaces.Geometry) *Decimate {
	return &Decimate{
		Geometry: geometry,
	}
}

// ValidateInputPointList validates the input point list.
//
// Parameters:
//   - points ([][]float64): The list of points to be validated.
//
// Returns:
//   - error: An error if the input point list is not valid.
//   - nil: If the input point list is valid.
func (d Decimate) ValidateInputPointList(points [][]float64) error {

	errorMsg := ""
	// size_points := len(points)
	// if size_points < 2 {
	// 	errorMsg += fmt.Sprintf("Length of point list must be greater than one and is %v, %v", size_points, points)
	// 	return errors.New(errorMsg)
	// }
	// return nil

	if len(points) > 1 {
		for i, point := range points {
			if len(point) != d.Geometry.Dimension() {
				errorMsg += fmt.Sprintf("All points must have the same dimension as the geometry. Point at position %v has dimension %v, but the geometry has dimension %v", i, len(point), d.Geometry.Dimension())
			}
		}
	}

	if errorMsg != "" {
		return errors.New(errorMsg)
	}

	return nil

}

// DouglasPeucker is a function that simplifies a list of points using the Douglas-Peucker algorithm.
//
// Parameters:
//   - points ([][]float64): The list of points to be simplified.
//   - threshold (float64): The threshold to be used in the simplification.
//
// Returns:
//   - [][]float64: The simplified list of points.
func (d Decimate) DouglasPeucker(points [][]float64, threshold float64) [][]float64 {

	errorMsg := d.ValidateInputPointList(points)

	if errorMsg != nil {
		panic(errorMsg)
	}

	if len(points) == 2 {
		return points
	}

	var distance float64
	distance_maximum := 0.0

	size_points := len(points)
	point_1 := primitives.NewPoint(points[0])
	point_2 := primitives.NewPoint(points[size_points-1])
	line := primitives.NewLine(point_1, point_2)

	var index int
	for i := range points {
		if i != 0 && i != size_points-1 {
			point_0 := primitives.NewPoint(points[i])
			distance = d.Geometry.DoubleAreaTriangle(point_0, line)
			if distance > distance_maximum {
				distance_maximum = distance
				index = i
			}
		}
	}

	point_0 := primitives.NewPoint(points[index])
	distance_maximum = d.Geometry.DistancePointLine(point_0, line)

	if distance_maximum < threshold {
		return [][]float64{points[0], points[size_points-1]}
	}
	left := d.DouglasPeucker(points[:index+1], threshold)
	right := d.DouglasPeucker(points[index:], threshold)
	return append(left[:len(left)-1], right...)

}
