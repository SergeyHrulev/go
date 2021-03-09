package main

import (
	"fmt"
)

func main()  {
	var number int32
	count := 0

	fmt.Scan(&number)

	for i := 1; int32(i) <= number && number < 1000000; i++ {
		if number % int32(i) == 0 {
			count++
			fmt.Printf("%d ", i)
			//fmt.Print(" ")
		}
	}
	if count == 2 {
		fmt.Printf("\n")
		fmt.Println("ACHTUNG")
	}
}