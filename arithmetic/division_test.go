package arithmetic

import "testing"

func TestDivide(t *testing.T) {
	result := Divide(10.0, 2.0)
	expected := 5.0
	if result != expected {
		t.Errorf("Divide(10.0, 2.0) = %f; want %f", result, expected)
	}
}
