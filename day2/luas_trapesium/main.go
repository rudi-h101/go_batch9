package main

import "fmt"

func main() {
	a := 10.5 // sisi atas
	b := 5.4 // sisi bawah
	t := 1.2 // tinggi
	
	L := 0.5 * (a + b) * t
	fmt.Println(L)
}