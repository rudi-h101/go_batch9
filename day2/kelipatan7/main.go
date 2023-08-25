package main

import "fmt"

func main() {
	N := 21

	if N < 7 && N % 7 == 0 {
		fmt.Println("Bilangan Kelipatan 7")
	} else {
		fmt.Println("Bilangan Bukan Kelipatan 7")
	}
}
