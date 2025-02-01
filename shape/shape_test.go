package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("test reactangle Area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}

		got := rectangle.Area()
		want := 200.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("test circle Area", func(t *testing.T) {
		circle := Circle{10}

		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
