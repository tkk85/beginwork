package main

import "fmt"

func main() {
	for i := 2; i <= 100; i++ {
		if (i == 2) || (i == 3) {
			fmt.Printf("The numbe %d is prime\n", i)
		} else {
			var count int = 0
			for j := 2; i <= (i - 1); j++ {
				flag := i % j
				if flag != 0 {
					count += 1
				}
				if count == (i - 2) {
					fmt.Printf("the number %d is prime\n", i)
				}
			}
		}
	}
}
