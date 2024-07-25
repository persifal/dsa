package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{5, 2, 3, 1}
	quickSort(arr)
	fmt.Printf("%d\n", arr)
}

func quickSort(arr []int) {
	n := len(arr)
	sort(arr, 0, n-1)
}

func sort(arr []int, left, right int) {
	if left < right {
		partitionIndex := partition(arr, left, right)
		sort(arr, left, partitionIndex-1)
		sort(arr, partitionIndex+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := takePivot(arr, left, right)
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	i++
	arr[i], arr[right] = arr[right], arr[i]

	return i
}

func takePivot(arr []int, left, right int) int {
	rnd := rand.Intn(right-left) + left
	arr[right], arr[rnd] = arr[rnd], arr[right]

	return arr[right]
}
