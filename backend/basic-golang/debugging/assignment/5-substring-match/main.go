package main

import "fmt"

func main() {
	/*
		Check whether the search string is exist in source string

		Sample Input/Output
		IsExistInSource("hello", "ll") -> True
		IsExistInSource("hello", "hel") -> True
		IsExistInSource("aaaa", "bb") -> False
	*/
	res := IsExistInSource("aaaa", "bb")
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := IsExistInSourceCorrect("hello", "ll")
	// fmt.Println(resCorrect)
}

func IsExistInSource(source, search string) bool {
	for startSource := 0; startSource < len(source); startSource++ {
		found := true
		idxSource := startSource

		for idxSearch := 0; idxSearch < len(search); idxSearch++ {
			if source[idxSource] == source[idxSearch] {
				found = true
				break
			}
			idxSearch++
		}
		if found {
			return true
		}
	}
	return false
}

func IsExistInSourceCorrect(source, search string) bool {
	for startSource := 0; startSource < len(source)-len(search)+1; startSource++ {
		found := true
		idxSource := startSource
		idxSearch := 0

		for idxSearch < len(search) {
			if source[idxSource] != search[idxSearch] {
				found = false
				break
			}
			idxSearch++
			idxSource++
		}
		if found {
			return true
		}
	}

	return false
}
