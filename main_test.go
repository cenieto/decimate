package main

import (
	"testing"
)


func TestReadPoints(t *testing.T) {
	result := readPoints("./fixtures/input-4d.json")
	expected := []Point4D{
		{1, 1, 0, 0}, {2, 2, 0, 0}, {3, 3, 0, 0}, {4, 4, 0, 0}, {5, 5, 0, 0}, {6, 6, 0, 0},
	}
	if len(result) != len(expected) {
		t.Errorf("readPoints(\"./fixtures/input-4d.json\") = %v; want %v", result, expected)
	}

	for index, value := range result{
		if value != expected[index]{
			t.Errorf("readPoints(\"./fixtures/input-4d.json\") = %v; want %v", result, expected)
		}
	}
}
