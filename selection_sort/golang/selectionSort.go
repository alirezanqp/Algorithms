package main

import "fmt"

func findSmallest(arr []int) int {
	smallest := arr[0]

	smallest_index := 0

	for i := 1; i < len(arr); i++ {
		if smallest > arr[i] {
			smallest = arr[i]
			smallest_index = i
		}
	}

	return smallest_index
}

func selectionSort(arr []int) []int {
	size := len(arr)
	newArray := make([]int, size)

	for i := 0; i < size; i++ {
		smallest := findSmallest(arr)
		newArray[i] = arr[smallest]

		arr = append(arr[:smallest], arr[smallest+1:]...)
	}

	return newArray
}

func main() {
	arr := []int{1, 5, 3, 2, 8, 6}
	fmt.Println(selectionSort(arr))
}