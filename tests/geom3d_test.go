package main

import (
	"decimator/pkg/geom3d"
	"reflect"
	"testing"
)

func TestCrossProduct3D(t *testing.T) {
	expected := []float64{0, 0, 1}
	v1 := geom3d.Vector3D{1, 0, 0}
	v2 := geom3d.Vector3D{0, 1, 0}
	result, _ := v1.CrossProduct(v2)

	if !reflect.DeepEqual(result.Components(), expected) {
		t.Errorf("v1.CrossProduct(v2) = %v; want %v", result.Components(), expected)
	}
}
