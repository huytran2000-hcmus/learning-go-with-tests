package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10, 10}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.g, want %.g", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{12, 6}, hasArea: 72},
		{name: "Circle", shape: &Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{12, 6}, hasArea: 36},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.hasArea {
				t.Errorf("%+v got %g, want %g", testCase.name, got, testCase.hasArea)
			}
		})
	}
}
