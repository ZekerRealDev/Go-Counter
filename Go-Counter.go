package main

import "fmt"

func main() {

	var number int

	fmt.Print("Enter the number up to which you want the count: ")
	fmt.Scanln(&number)

	var i int
	for i = 1; i <= number; i++ {

		fmt.Println(i)

	}

}
