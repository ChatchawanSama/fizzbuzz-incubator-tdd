package level3

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when input is 1", func(t *testing.T) {
		result := FizzBuzz(1)
		expected := "1"
		if result != expected {
			t.Errorf("FizzBuzz(1) = %s, want %s", result, expected)
		}
	})
}
