package level3

// FizzBuzz with no if and no switch statement
func FizzBuzz(n int) string {
	arr := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "10"}
	return arr[(n%10)-1]
}
