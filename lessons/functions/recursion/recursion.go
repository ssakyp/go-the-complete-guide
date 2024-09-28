package recursion

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
	
	// res := 1

	// for i := 1; i <= number; i++ {
	// 	res = res * i
	// }

	// return res
}

// if a function call itself we use recursion
// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120
