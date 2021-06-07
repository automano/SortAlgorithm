package main

import (
	"fmt"

	"github.com/automano/SortAlgorithm/mysort"
)

func main() {
	// 切片
	data := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	fmt.Println("Before sorting: data = ", data)
	// mysort.BubbleSort(data)
	// mysort.SelectionSort(data)
	// mysort.InsertionSort(data)
	// mysort.ShellSort(data)
	// data = mysort.MergeSort(data)
	// mysort.QuickSort(data)
	mysort.RandomQuickSort(data)
	fmt.Println("After sorting: data = ", data)
}
