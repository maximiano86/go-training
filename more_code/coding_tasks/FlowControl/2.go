package main

import "fmt"

func main() {
	for i := 0; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
}
