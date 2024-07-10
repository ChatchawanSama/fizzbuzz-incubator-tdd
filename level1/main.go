package level1

import "fmt"

func FizzBuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprint(n)
}
