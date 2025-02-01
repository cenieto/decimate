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
