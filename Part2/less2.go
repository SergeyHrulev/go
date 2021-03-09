package main

import "fmt"

func main()  {
	var (
		a float64
		b float64
		c float64
	)

	fmt.Scan(&a, &b, &c)
	if a == 0.00 {
		if b != 0.00 {
			fmt.Println("два корня")
		} else if b == 0 {
			fmt.Println("один корень")
		}
	}
	if (b * b) - (4 * a * c) < 0 {
		fmt.Println("корней нет")
	} else if (b * b) - (4 * a * c) == 0 {
		fmt.Println("один корень")
	} else if (b * b) - (4 * a * c) > 0 {
		fmt.Println("два корня")
		return
	}
}
