package arithmetic

import "testing"

func TestMultiply(t *testing.T) {
    result := Multiply(4.0, 2.5)
    expected := 10.0
    if result != expected {
        t.Errorf("Multiply(4.0, 2.5) = %f; want %f", result, expected)
    }
}
