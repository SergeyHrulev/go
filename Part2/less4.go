package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main()  {
	var (
		a string
		b string
	)

	for {
		fmt.Scanf("%s\n", &a)
		fmt.Scanf("%s\n", &b)
		if utf8.RuneCountInString(a) < 8 {
			fmt.Println("Слишком короткий пароль!")
		} else if strings.Contains(a, "123") || strings.Contains(a, "qwe") {
			fmt.Println("Слишком простой пароль!")
		} else if a != b {
			fmt.Println("Введенные пароли различаются!")
		} else {
			fmt.Println("Ну наконец-то!")
			return
		}
	}
}