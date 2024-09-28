package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	double := doubleNumbers(&numbers)
	fmt.Println(double)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, double(v))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
