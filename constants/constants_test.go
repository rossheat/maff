package constants

import "testing"

func TestPi(t *testing.T) {
	expected := 3.14159
	if Pi != expected {
		t.Errorf("Pi = %f; want %f", Pi, expected)
	}
}
