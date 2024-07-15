package level2

import "fmt"

// FizzBuzz With One If No Switch
func FizzBuzz(n int) string {
	arr := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	if n >= 1 || n <= 15 {
		return fmt.Sprint(arr[(n%100)-1])
	}
	return ""
}
