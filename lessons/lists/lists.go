package main


/// different type of data grouped into one
// type Product struct {
//   title string
//   id string
//   price float64
// }

func main() {
  var productNames [4]string = {"A Book"}
  // productNames = [4]string{"A Book"}
  productNames[2] = "A Carpet"
  prices := [4]float64{10.99, 9.99, 45.99, 20.0}
  fmt.Println(prices) // [10.99 9.99 45.99 20.0]
  fmt.Println(productNames) // [A Book   A Carpet  ]

  fmt.Println(prices[2])  // 45.99

  // creating slices => 1st value included, 2nd value excluded
  featuredPrices := prices[1:3]    // [9.99 45.99]
  featuredPrices = prices[1:]    // [9.99 45.99 20.]
  featuredPrices[0] = 199.99
  highlightedPrices := featuredPrices[:1]    // [9.99] => 199.99
  fmt.Println(len(featuredPrices), cap(featuredPrices)) // 3, 3
  fmt.Println(len(highglightedPrices), cap(highglightedPrices)) // 1 3

  highlightedPrice = highlightedPrices[:3]
  fmt.Println(len(highglightedPrices), cap(highglightedPrices)) // 3 3
}
