package main

import "fmt"

func main() {
	var slice []int
	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	for _, integer := range slice {
		if integer%2 == 0 {
			fmt.Println(integer, "is even")
		} else {
			fmt.Println(integer, "is odd")
		}
	}
}
