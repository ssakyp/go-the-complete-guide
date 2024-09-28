package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	double := transformNumbers(&numbers, double)
	tripple := transformNumbers(&double, triple)
	fmt.Println(double)
	fmt.Println(tripple)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
