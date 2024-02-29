package main

import (
	"log"

	"github.com/highborn/playground/problems"
)

func main() {
	array := []int{10, 7, 8, 9, 1, 5}
	array2 := []int{10, 7, 8, 9, 1, 5}
	// array := []int{0, 0, 0}
	// array2 := []int{0, 0, 0}
	// array := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	// array := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	problems.ArrayRankTransform(array)
	log.Println(array)
	array2 = problems.ArrayRankTransform2(array2)
	log.Println(array2)
}
