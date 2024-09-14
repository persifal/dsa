package boyer_moore

func boyerMoore(arr []int) int {
	candidate := findCandidate(arr)
	if isDominante(arr, candidate) {
		return candidate
	} else {
		return -1
	}
}

func isDominante(arr []int, candidate int) bool {
	cnt := 0
	for _, v := range arr {
		if v == candidate {
			cnt++
		}
	}

	return cnt > len(arr)/2
}

func findCandidate(arr []int) int {
	candidate := 0
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if cnt == 0 {
			candidate = arr[i]
		}

		if arr[i] == candidate {
			cnt++
		} else {
			cnt--
		}
	}

	return candidate
}
