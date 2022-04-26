// Reverse

package main

import "fmt"

func Reverse(st []string, depth int) string {
	str := ""
	for i := 8; i >= 0; i-- {
		str += st[i]
		fmt.Println(st[i])
	}

	return str

	/*
		i = 1
		str = str + st[0] = "A"

		i = 2
		str = str + st[1] = "I"

		i = 3
		str = str + st[2] = "S"

		i = 4
		str = str + st[3] = "E"

		i = 5
		str = str + st[4] = "N"

		i = 6
		str = str + st[5] = "O"

		i = 7
		str = str + st[6] = "D"

		i = 8
		str = str + st[7] = "N"

	*/
}

func main() {
	str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
	s := Reverse(str, len(str)-1)
	fmt.Println(s)
}
