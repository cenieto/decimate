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
	// "gonum.org/v1/gonum/mat"
	"fmt"
	"github.com/cenieto/decimate/pkg/interfaces"
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

func (d Decimate) DouglasPeucker(points [][]float64) [][]float64 {
	// var distance float64
	// var distance_maximum float64
	// distance_maximum = 0.0

	errorMsg := ""
	size_points := len(points)
	if size_points < 2 {
		errorMsg += "Length of point list must be greater than one\n"
	}
	
	if len(points) > 1 {
		for i, point := range points {
			if len(point) != d.Geometry.Dimension() {
				errorMsg += fmt.Sprintf("All points must have the same dimension as the geometry. Point at position %v has dimension %v, but the geometry has dimension %v", i, len(point), d.Geometry.Dimension())
			}
		}
	}

	if errorMsg != "" {
		panic(errorMsg)
	}

	// // TODO add test to check return when only to points are given
	// if size_points == 2 {
	// 	return points
	// }

	// dimension := points[0].Dimension()
	// if dimension != 2 {
	// 	geometry = interfaces.NewGeometry(dimension)
	// }

	// var index int
	// for i := 1; i < size_points; i++ {
	// 	distance = PerpendicularDistance(points[i], points[0], points[size_points-1])
	// 	if distance > distance_maximum {
	// 		index = i
	// 		distance_maximum = distance
	// 	}
	// }

	return points
}
