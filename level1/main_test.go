package level1

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when input is 1", func(t *testing.T) {
		result := FizzBuzz(1)
		expected := "1"
		if result != expected {
			t.Errorf("FizzBuzz(1) = %s; want %s", result, expected)
		}
	})

	t.Run("should return 2 when input is 2", func(t *testing.T) {
		result := FizzBuzz(2)
		expected := "2"
		if result != expected {
			t.Errorf("FizzBuzz(2) = %s; want %s", result, expected)
		}
	})

	t.Run("should return Fizz when input is 3", func(t *testing.T) {
		result := FizzBuzz(3)
		expected := "Fizz"
		if result != expected {
			t.Errorf("FizzBuzz(3) = %s; want %s", result, expected)
		}
	})

	t.Run("should return 4 when input is 4", func(t *testing.T) {
		result := FizzBuzz(4)
		expected := "4"
		if result != expected {
			t.Errorf("FizzBuzz(4) = %s; want %s", result, expected)
		}
	})

	t.Run("should return Buzz when input is 5", func(t *testing.T) {
		result := FizzBuzz(5)
		expected := "Buzz"
		if result != expected {
			t.Errorf("FizzBuzz(5) = %s; want %s", result, expected)
		}
	})

	t.Run("should return Fizz when input is 6", func(t *testing.T) {
		result := FizzBuzz(6)
		expected := "Fizz"
		if result != expected {
			t.Errorf("FizzBuzz(6) = %s; want %s", result, expected)
		}
	})

}
