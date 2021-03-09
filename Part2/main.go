package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	var logo string
	var price int

	fmt.Scan(&logo)

	price = utf8.RuneCountInString(logo) * 23

	if price < 100 {
		fmt.Println(price, " Коп.")
	}
	rub := price / 100
	kop := price % 100
	fmt.Printf("%d р. %d коп.\n", rub, kop)
}
