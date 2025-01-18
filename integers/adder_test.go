package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errof("expected '%d' but got '%d'", expected, sum)
	}
}
