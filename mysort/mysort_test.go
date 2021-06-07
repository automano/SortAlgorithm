package mysort_test

import (
	"testing"

	"github.com/automano/SortAlgorithm/mysort"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

// IsSorted reports whether data is sorted.
func IsSorted(data []int) bool {
	n := len(data)
	for i := n - 1; i > 0; i-- {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	data := ints
	mysort.BubbleSort(data[0:])
	if !IsSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestInsertionSort(t *testing.T) {
	data := ints
	mysort.InsertionSort(data[0:])
	if !IsSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestMergeSort(t *testing.T) {
	data := ints[0:]
	data = mysort.MergeSort(data)
	if !IsSorted(data) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestQuickSort(t *testing.T) {
	data := ints
	mysort.QuickSort(data[0:])
	if !IsSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestRandomQuickSort(t *testing.T) {
	data := ints
	mysort.RandomQuickSort(data[0:])
	if !IsSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestSelectionSort(t *testing.T) {
	data := ints
	mysort.SelectionSort(data[0:])
	if !IsSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestShellSort(t *testing.T) {
	data := ints
	mysort.ShellSort(data[0:])
	if !IsSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func BenchmarkBubbleSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		mysort.BubbleSort(data)
		b.StopTimer()
	}
}

func BenchmarkInsertionSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		mysort.InsertionSort(data)
		b.StopTimer()
	}
}

func BenchmarkMergeSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		_ = mysort.MergeSort(data)
		b.StopTimer()
	}
}

func BenchmarkQuickSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		mysort.QuickSort(data)
		b.StopTimer()
	}
}

func BenchmarkRandomQuickSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		mysort.RandomQuickSort(data)
		b.StopTimer()
	}
}

func BenchmarkSelectionSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		mysort.Sele(data)
		b.StopTimer()
	}
}

func BenchmarkRandomQuickSortInt1K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		mysort.RandomQuickSort(data)
		b.StopTimer()
	}
}
