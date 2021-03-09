package main

import (
	"fmt"
)

func main()  {
	var (
		countElements int
		switcher bool = true
		sum int
		element int
	)

	fmt.Scan(&countElements)

	for i := 1; i <= countElements; i++ {
		fmt.Scan(&element)
		if switcher {
			sum += element
		} else if !switcher {
			sum -= element
		}
		switcher = !switcher
	}
	fmt.Println(sum)
}