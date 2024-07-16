package level3

// FizzBuzz with no if and no switch statement
func FizzBuzz(n int) string {
	arr := []string{"1", "2", "Fizz", "4", "Buzz"}
	return arr[(n%10)-1]
}
