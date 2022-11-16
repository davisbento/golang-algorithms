package algorithms

import (
	"fmt"
	"math"
)

func BinarySearchTree() {
	fmt.Println("------------- Binary Search Tree -------------")

	// Node represents a node in a binary search tree.
	data := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 9

	nLoops := 0

	// O(n)
	for _, x := range data {
		nLoops++
		if x == target {
			fmt.Printf("Found %d \n", target)
			break
		}
	}

	fmt.Printf("Number of loops O(n): %d \n", nLoops)

	nLoops = 0

	// O(log n)
	// binary search
	low := 0
	high := len(data) - 1

	for low <= high {
		nLoops++
		mid := int(math.Floor(float64(low+high) / 2))

		if data[mid] == target {
			fmt.Printf("Found %d \n", target)
			break
		}

		if data[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Printf("Number of loops O(log n): %d \n", nLoops)
	fmt.Println("------------- End Binary Search Tree -------------")
}
