package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func newProduct(title string, id string, price float64) Product {
	return Product{title, id, price}
}

func main() {

	hobbies := [3]string{"first", "second", "third"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//new := hobbies[0:2]
	new := hobbies[:2]
	//new := [2]string{hobbies[0], hobbies[1]}
	fmt.Println(new)

	new = new[1:3]
	fmt.Println(new)

	slice := []string{"one", "two", "three"}
	slice[1] = "double"
	slice = append(slice, "fourth")

	firstProd := newProduct("prod", "first", 2.22)
	secondProd := newProduct("prod", "second", 3.25)
	prodList := []Product{firstProd, secondProd}
	thirdProd := newProduct("prod", "third", 4.55)
	prodList = append(prodList, thirdProd)
	fmt.Println(prodList)

}
