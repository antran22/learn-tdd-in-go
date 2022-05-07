package structs

import (
	"testing"
)

const epsilon = 1e-6

func compareFloat(expected, got float64) bool {
	diff := expected - got
	if diff > epsilon || diff < -epsilon {
		return true
	}
	return false
}

type perimeterTestCase struct {
	name         string
	shape        Shape
	hasPerimeter float64
}

func TestPerimeter(t *testing.T) {
	testCases := []perimeterTestCase{
		{
			name:         "Rectangle",
			shape:        Rectangle{4.3, 6.1},
			hasPerimeter: 20.8,
		},
		{
			name:         "Circle",
			shape:        Circle{2.4},
			hasPerimeter: 15.0796447372,
		},
		{
			name:         "RightTriangle",
			shape:        RightTriangle{12, 6},
			hasPerimeter: 31.416407865,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Perimeter()
			if compareFloat(testCase.hasPerimeter, got) {
				t.Errorf("Error on test case %#v, hasPerimeter %g, got %g", testCase, testCase.hasPerimeter, got)
			}
		})
	}
}

type areaTestCase struct {
	name    string
	shape   Shape
	hasArea float64
}

func TestArea(t *testing.T) {
	testCases := []areaTestCase{
		{
			name:    "Rectangle",
			shape:   Rectangle{4.3, 6.1},
			hasArea: 26.23,
		},
		{
			name:    "Circle",
			shape:   Circle{3.2},
			hasArea: 32.1699087728,
		},
		{
			name:    "RightTriangle",
			shape:   RightTriangle{12, 6},
			hasArea: 36,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if compareFloat(testCase.hasArea, got) {
				t.Errorf("Error on test case %#v, hasPerimeter %g, got %g", testCase, testCase.hasArea, got)
			}
		})
	}
}
