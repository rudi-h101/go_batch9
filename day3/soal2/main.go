package main

import "fmt"

func Mapping(slice []string) map[string]int {
	result := map[string]int{}

	for _, element := range slice {
		if(result[element] == 0){
			result[element] = 1
		}else{
			result[element] += 1
		}
	}

	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	fmt.Println(Mapping([]string{}))
}