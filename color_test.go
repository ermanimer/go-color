package color

import "testing"

func TestWith(t *testing.T) {
	expected := escCode("\x1b[38;5;7;m\x1b[48;5;1;m")
	actual := White.With(Red)
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}
