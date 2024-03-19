package main

import "fmt"

func main() {
	// Define the two input arrays
	array1 := []int{1, 2, 3}
	array2 := []int{4, 5, 6}

	// Calculate the sum of corresponding elements and store in a third array
	var sumArray []int
	for i := 0; i < len(array1); i++ {
		sumArray = append(sumArray, array1[i]+array2[i])
	}

	// Print the individual sums in the third array
	fmt.Println("Sum Array:", sumArray)
}
