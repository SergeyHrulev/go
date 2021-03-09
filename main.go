package main

import (
	"fmt"
)

func main() {

	var (
		a int
		b int
		c int
		d int
	)

	fmt.Scan(&a, &b, &c, &d)
	if c - a == 1 && d - b == 2 {
		fmt.Println( "ДА")
		return
	} else if c - a == 2 && d - b == 1 {
		fmt.Println( "ДА")
		return
	} else if c - a == -1 &&  d - b == -2 {
		fmt.Println( "ДА")
		return
	} else if c - a == -2 && d - b == -1 {
		fmt.Println( "ДА")
		return
	}
	fmt.Println("НЕТ")
}