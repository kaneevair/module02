package main

import "fmt"

func main() {
	var A *int
	var B = 10

	A = &B
	fmt.Println(*A)

	*A = 15
	fmt.Println(B)

}
