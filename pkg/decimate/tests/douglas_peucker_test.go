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
	"testing"
	"github.com/cenieto/decimate/pkg/geom2d"
)

// TestDouglasPeuckerInputCheck tests the input check of the DouglasPeucker function.
// It checks if the function panics when the input has only one element.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestDouglasPeuckerInputCheckOnePoint(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
	}

	geometry := geom2d.NewEuclid()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DouglasPeucker() did not panic when input had only one element")
		}
	}()

	new_points := geometry.Decimate.DouglasPeucker(points)

	fmt.Println(new_points)
}

// TestDouglasPeuckerInputCheckDifferntDimensions tests the input check of the DouglasPeucker function.
// It checks if the function panics when the input has points with dimension different to the geometry.
//
// Parameters:
//   - t (*testing.T): A testing object used to run tests and check for failures.
//
// Returns:
//   - None
func TestDouglasPeuckerInputCheckDifferntDimensions(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{1.0, 2.0, 3.0},
	}

	geometry := geom2d.NewEuclid()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DouglasPeucker() did not panic when input had points with dimension different to geometry")
		}
	}()

	new_points := geometry.Decimate.DouglasPeucker(points)

	fmt.Println(new_points)
}
