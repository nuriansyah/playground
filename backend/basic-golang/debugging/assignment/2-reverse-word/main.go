package main

import "fmt"

func main() {
	/*
		Reverse Word:
		Example: halo -> olah
	*/
	word := "halo"
	res := ReverseWord(word)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ReverseWordCorrect(arr)
	// fmt.Println(resCorrect)
}

func ReverseWord(word string) string {
	n := len(word)
	temp := []byte(word)

	for i := 0; i <= n; i++ {
		left := i
		right := n - i - 1
		temp[left], temp[right] = temp[right], temp[left]
	}

	return string(temp)
}

func ReverseWordCorrect(word string) string {
	if len(word) == 0 {
		return ""
	}

	if len(word) == 1 {
		return word
	}

	return word[len(word)-1:] + ReverseWordCorrect(word[:len(word)-1])
}
