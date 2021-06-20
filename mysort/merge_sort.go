package mysort

// 归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。

// 算法描述
// - 把长度为n的输入序列分成两个长度为n/2的子序列；
// - 对这两个子序列分别采用归并排序；
// - 将两个排序好的子序列合并成一个最终的排序序列

// 归并排序是一种稳定的排序方法。和选择排序一样，归并排序的性能不受输入数据的影响，但表现比选择排序好的多，因为始终都是O(nlogn）的时间复杂度。代价是需要额外的内存空间。

// MergeSort 归并排序
// Merge sort is an effective sort algorithm based on merge operation. It is a very typical application of the Divide and Conquer method.
func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) >> 1
	left, right := MergeSort(data[:middle]), MergeSort(data[middle:])
	merge := func(left, right []int) []int {
		result := make([]int, len(left)+len(right))
		for i := 0; len(left) > 0 || len(right) > 0; i++ {
			if len(left) > 0 && len(right) > 0 {
				if left[0] < right[0] {
					result[i] = left[0]
					left = left[1:]
				} else {
					result[i] = right[0]
					right = right[1:]
				}
			} else if len(left) > 0 {
				result[i] = left[0]
				left = left[1:]
			} else if len(right) > 0 {
				result[i] = right[0]
				right = right[1:]
			}
		}
		return result
	}
	return merge(left, right)
}
