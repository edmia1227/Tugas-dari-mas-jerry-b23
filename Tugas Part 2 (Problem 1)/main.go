package main

import "fmt"

func main() {

	n := 5

	for i := 1; i <= n; i++ {

		for j := 1; j <= n-i; j++ {

			fmt.Print(" ")

		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")

		}
		fmt.Println()

	}
}
