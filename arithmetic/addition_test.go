package arithmetic

import "testing"

func TestAdd(t *testing.T) {
    result := Add(5.0, 3.0)
    expected := 8.0
    if result != expected {
        t.Errorf("Add(5.0, 3.0) = %f; want %f", result, expected)
    }
}
