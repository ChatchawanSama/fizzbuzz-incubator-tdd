package level2

import "fmt"

//FizzBuzz With One If No Switch
func FizzBuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprint(n)
}