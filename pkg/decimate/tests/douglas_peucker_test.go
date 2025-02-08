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
package tests

import (
	"fmt"
	"github.com/cenieto/decimate/pkg/geom2d"
	"testing"
)

// TestValidateInputPointListOnePoint tests the input check of the ValidateInputPointList function.
// It checks if the function panics when the input has only one element.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestValidateInputPointListOnePoint(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
	}

	geometry := geom2d.NewEuclid()

	errorMsg := geometry.Decimate.ValidateInputPointList(points)
	if errorMsg != nil {
		return errors.New("Length of point list must be greater than one")
	}

}

// TestValidateInputPointListDifferntDimensions tests the input check of the ValidateInputPointList function.
// It checks if the function panics when the input has points with different dimensions.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestValidateInputPointListDifferntDimensions(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{1.0, 2.0, 3.0},
	}

	geometry := geom2d.NewEuclid()

	errorMsg := geometry.Decimate.ValidateInputPointList(points)

	if errorMsg != nil {
		return errors.New("All points must have the same dimension as the geometry. Point at position %v has dimension %v, but the geometry has dimension %v", i, len(point), d.Geometry.Dimension())
	}

}
