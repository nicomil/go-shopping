package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type prod struct {
	name  string
	price float32
}

var numbers []int
var products = []prod{
	{"Water", 5},
	{"Bread", 3},
	{"Cola", 4},
}
var cart []prod

var wallet float32 = 10

func main() {
	minPrice := minPrice(products)
	reader := bufio.NewReader(os.Stdin)
	for {
		showProducts(products)

		product := products[choseProduct(reader)]

		// n, _ := strconv.ParseInt(slice_num, 10, 0) // alternative to Atoi
		if product.price <= wallet {
			cart = append(cart, product)
			wallet -= product.price
		}

		fmt.Println(cart)
		fmt.Print("Another one?(yes/no) ")
		resp, _ := reader.ReadString('\n')
		// fmt.Printf("%#v", resp)
		// fmt.Printf("%#v", "no")

		if strings.Compare(resp, "no\n") == 0 || minPrice > wallet {
			fmt.Println("Bye")
			break
		}
	}
	fmt.Println("Your cart", cart)
	fmt.Println("Your wallet", wallet)
}

func minPrice(arr []prod) float32 {
	var min float32 = arr[0].price
	for _, el := range arr {
		if min > el.price {
			min = el.price
		}
	}
	return min
}

func showProducts(products []prod) {
	for i, product := range products {
		fmt.Println("#", i+1)
		fmt.Println("Prod: ", product.name)
		fmt.Println("Price: ", product.price)
	}
}

func choseProduct(reader *bufio.Reader) int {
	fmt.Println("Which product want to buy?")
	fmt.Print("Enter number: ")
	num, _ := reader.ReadString('\n')
	slice_num := strings.TrimSuffix(num, "\n")
	n, _ := strconv.Atoi(slice_num)

	return n
}
