package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	pipe(arr)
}

// cal squre of number
func genPowTwo(nums []int, out chan int) {
	go func() {
		for _, n := range nums {
			out <- n * n
		}
		close(out)
	}()
}

// sum number recive from channel
func sum(in <-chan int, result chan int) {
	var res int
	go func() {
		for n := range in {
			res = res + n
		}
		result <- res
		close(result)
	}()
}

// calulate sum of ech number by power of 2
func pipe(nums []int) {
	out := make(chan int)
	result := make(chan int)
	go genPowTwo(nums, out)
	go sum(out, result)

	res := <-result
	fmt.Println(res)
}
