package main

import (
	"decimator/pkg/geom2d"
	"fmt"
	"reflect"
	"testing"
)

func TestCrossProduct2D(t *testing.T) {
	expected := []float64{0, 0, 1}
	v1 := geom2d.NewVector(1, 0)
	v2 := geom2d.NewVector(0, 1)
	result, _ := v1.CrossProduct(v2)

	if !reflect.DeepEqual(result.Components(), expected) {
		t.Errorf("v1.CrossProduct(v2) = %v; want %v", result.Components(), expected)
	}
}

func TestSegment(t *testing.T) {
	p1 := geom2d.NewPoint(1, 0)
	p2 := geom2d.NewPoint(0, 1)
	s1 := geom2d.NewSegment(p1, p2)

	fmt.Println(s1)
}
