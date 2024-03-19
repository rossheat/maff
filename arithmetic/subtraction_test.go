package arithmetic

import "testing"

func TestSubtract(t *testing.T) {
    result := Subtract(10.0, 4.0)
    expected := 6.0
    if result != expected {
        t.Errorf("Subtract(10.0, 4.0) = %f; want %f", result, expected)
    }
}
