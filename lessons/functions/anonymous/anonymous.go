package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	double := creatTransformer(2)
	triple := creatTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubledNumbers := transformNumbers(&numbers, double)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}

	return dNumbers
}

func creatTransformer(factor int) func(int) int {
	return func(number int) int{
		return number * factor
	}
}