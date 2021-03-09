package main

import "fmt"

func main()  {
	var (
		a float64 = 0.00
		counter int
		min float64 = 140
		max float64 = 100
	)
	for a >= 0{
		fmt.Scan(&a)
		if a >= 100 && a <= 140 {
			counter++
			if a <= min {
				min = a
			}
			if max <= a {
				max = a
			}
		}
	}
	fmt.Println(counter)
	fmt.Printf("%.1f %.1f\n", min, max)
}