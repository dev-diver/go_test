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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("test reactangle Area", func(t *testing.T) {

		rectangle := Rectangle{10.0, 20.0}
		want := 200.0
		checkArea(t, rectangle, want)

	})

	t.Run("test circle Area", func(t *testing.T) {

		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)

	})

}
