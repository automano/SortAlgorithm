package mysort

func RadixSort(data []int) {
	largestNum := findLargestNum(data)
	size := len(data)
	significantDigit := 1
	semiSorted := make([]int, size)

	for largestNum/significantDigit > 0 {
		bucket := [10]int{0}

		for i := 0; i < size; i++ {
			bucket[(data[i]/significantDigit)%10]++
		}

		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		for i := size - 1; i >= 0; i-- {
			bucket[(data[i]/significantDigit)%10]--
			semiSorted[bucket[(data[i]/significantDigit)%10]] = data[i]
		}

		for i := 0; i < size; i++ {
			data[i] = semiSorted[i]
		}

		significantDigit *= 10
	}
}

func findLargestNum(data []int) int {
	largestNum := 0

	for i := 0; i < len(data); i++ {
		if data[i] > largestNum {
			largestNum = data[i]
		}
	}
	return largestNum
}
