package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	result := []string{}

	temp := map[string]int{}

	newArray := append(arrayA, arrayB...)

	for _, element := range newArray {
		if(temp[element] == 0){
			result = append(result, element)
			temp[element] = 1
		}
	}

	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{ "king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa","law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{ "hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}

