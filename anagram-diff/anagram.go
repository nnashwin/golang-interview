package main

import (
	"fmt"
	"math"
)

func CreateMap(str string) (strMap map[string]float64) {
	strMap = make(map[string]float64)
	for idx, _ := range str {
		char := string(str[idx])
		if strMap[char] != 0 {
			strMap[char] += 1
		} else {
			strMap[char] = 1
		}
	}
	return
}

func StrDiff(str1 string, str2 string) (count float64) {
	hash1 := CreateMap(str1)
	hash2 := CreateMap(str2)

	for char, _ := range hash1 {
		if hash2[char] == 0 {
			count += hash1[char]
		} else {
			count += math.Max(0, hash1[char]-hash2[char])
		}
	}

	for char, _ := range hash2 {
		if hash1[char] == 0 {
			count += hash1[char]
		} else {
			count += math.Max(0, hash2[char]-hash1[char])
		}
	}

	return
}

func main() {
	count := StrDiff("tyler", "trever")
	fmt.Println(count)
}
