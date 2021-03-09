package main

import (
	"fmt"
)

func main()  {
	var (
		n int
		m int
		arr1 [10]string
		arr2 [10]string
	)

	fmt.Scan(&n, &m)

	// A,H,I Задачи

	for i := 0; i <= n - 1; i++ {
		//fmt.Println("Индекс i: ", i)
		fmt.Scan(&arr1[i])
		fmt.Scanf("%q", &name)
	}
	for j := 0; j <= m - 1; j++ {
		//fmt.Println("Индекс j: ", j)
		fmt.Scan(&arr2[j])
	}
	fmt.Println(arr1)
	//for _, val := range arr1 {
	//	fmt.Println(val)
	//}
	//for _, val1 := range arr1 {
	//	for _, val2 := range arr2{
	//		if strings.Contains(val1, val2) {
	//			fmt.Println("ЕСТЬ")
	//			break
	//		} else {
	//			fmt.Println("НЕТ В НАЛИЧИИ")
	//			break
	//		}
	//	}
	//}
}
