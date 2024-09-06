package quickselect

func quickSelect(arr []int, index int) int {
	n := len(arr)
	return doSelect(arr, 0, n-1, index)
}

func doSelect(arr []int, left, right, index int) int {
	partitionIndex := partition(arr, left, right)
	if partitionIndex == index-1 {
		return arr[index]
	} else if partitionIndex < index-1 {
		return doSelect(arr, partitionIndex+1, right, index)
	} else {
		return doSelect(arr, left, partitionIndex-1, index)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	pivotLocation := left
	for j := left; j <= right; j++ {
		if arr[j] < pivot {
			arr[pivotLocation], arr[j] = arr[j], arr[pivotLocation]
			pivotLocation++
		}
	}

	arr[pivotLocation], arr[right] = arr[right], arr[pivotLocation]

	return pivotLocation
}

// func randomizedPivot(arr []int, left, right int) int {
// 	rnd := rand.Intn(right-left) + left
// 	arr[right], arr[rnd] = arr[rnd], arr[right]

// 	return arr[right]
// }
