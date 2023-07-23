package main

import "fmt"

func main() {
	sum := 1
	// goではwhileはforだけでいける
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}