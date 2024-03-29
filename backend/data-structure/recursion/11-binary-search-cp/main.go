package main

import "fmt"

func main() {
	numList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := int64(9)
	fmt.Println(BinarySearch(numList, key))
}

//Recursive Binary Search
func BinarySearch(numList []int64, key int64) int {
	low := 0
	high := len(numList) - 1

	if low <= high {
		// TODO: answer here
		mid := low + (high-low)/2
		//if binary search is found, return 1
		if numList[mid] == key {
			return 1
		}
		//if binary search is not found, return 0
		if numList[mid] > key {
			return BinarySearch(numList[:mid], key)
		} else {
			return BinarySearch(numList[mid+1:], key)
		}
	}
	return 0
}
