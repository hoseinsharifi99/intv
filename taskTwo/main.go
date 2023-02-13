package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 4, 5, 3, 2, 1}
	fmt.Println(FindOneRepeted(arr))
}

// find number dont repeted twitce
func FindOneRepeted(arr []int) int {
	var res int
	for _, value := range arr {
		//xor
		res = res ^ value
	}

	return res
}
