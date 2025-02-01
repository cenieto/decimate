// Copyright 2025 César Nieto Sánchez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
