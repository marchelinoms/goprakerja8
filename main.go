package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	var merged []string
	nameMap := make(map[string]bool)

	for _, name := range arrayA {
		nameMap[name] = true
		merged = append(merged, name)
	}

	for _, name := range arrayB {
		if _, exists := nameMap[name]; !exists {
			nameMap[name] = true
			merged = append(merged, name)
		}
	}

	return merged
}

func Mapping(slice []string) map[string]int {
	result := make(map[string]int)
	for _, kata := range slice {
		result[kata]++
	}
	return result
}

func munculSekali(input string) []int {
	numCount := make(map[int]int)
	uniqueNumber := make([]int, 0)

	for _, angka := range input {
		num := int(angka - '0')
		numCount[num]++
	}

	for num, count := range numCount {
		if count == 1 {
			uniqueNumber = append(uniqueNumber, num)
		}
	}
	return uniqueNumber
}

func main() {
	//Test Cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{}))
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
