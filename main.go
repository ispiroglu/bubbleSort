package main

import "fmt"

func main() {
	arr := []int{3, -2, 99, 6, -55, 8, 44, 2}

	fmt.Println("Array before sorting.")
	printArray(arr)

	fmt.Println()
	fmt.Println()
	bubbleSort(arr)

	fmt.Println("Array after sorting.")
	printArray(arr)

}

func bubbleSort(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if isGreater(arr, j) {
				swapInGo(arr, j)
			}
		}
	}
}

func isGreater(arr []int, j int) bool {
	if arr[j] > arr[j+1] {
		return true
	}

	return false
}

func swapInGo(arr []int, j int) {
	arr[j], arr[j+1] = arr[j+1], arr[j]
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
