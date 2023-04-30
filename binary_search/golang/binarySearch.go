package main

import "fmt"

func binarySearchIterative(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {

		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}

		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarySearchRecursive(list []int, item int, low int, high int) int {
    if low > high {
        return -1
    }

    mid := (low + high) / 2
    guess := list[mid]

    if guess == item {
        return mid
    } else if guess < item {
        return binarySearchRecursive(list, item, mid+1, high)
    } else {
        return binarySearchRecursive(list, item, low, mid-1)
    }
}

func main() {
	fmt.Println(binarySearchIterative([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(binarySearchIterative([]int{1, 2, 3, 4, 5}, 6))

	fmt.Println(binarySearchRecursive([]int{1, 2, 3, 4, 5}, 3, 0, 4))
    	fmt.Println(binarySearchRecursive([]int{1, 2, 3, 4, 5}, 6, 0, 4))
}
