package main

import (
	"fmt"
	"slices"
	"strconv"
)

func munculSekali(angka string) []int {
	
	result := []int{}
	temp := map[int]int{}
	
	for _, char := range angka {
		char, _ := strconv.Atoi(string(char))

		if (temp[char] == 0 && char != 0) {
			result = append(result, char)
			temp[char] = 1
		} else {
			// remove exsisting characters in main array
			contains := slices.Contains(result, char)
			if(contains) {
				charIndex := slices.Index(result, char)

				result = slices.Delete(result, charIndex, charIndex + 1)
			}
		}
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872564"))
}