package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	var a string

	for {
		fmt.Scanf("%s\n", &a)
		if utf8.RuneCountInString(a) > 0 {
			fmt.Println(a)
			a = ""
		} else if utf8.RuneCountInString(a) <= 0 {
			return
		}

	}
}
