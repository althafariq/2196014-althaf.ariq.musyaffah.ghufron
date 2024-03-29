// Golang program to count digits of given number
// using recursion

package main

import "fmt"

var count int = 0

//function to count digits
func CountDigits(num int) int {

	if num > 0 {
		count++
		//fmt.Println(num / 10)
		CountDigits(num / 10)

	}
	return count

}

func main() {
	var num int = 0
	var result int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	result = CountDigits(num)
	fmt.Printf("Count of digits is: %d\n", result)
}
