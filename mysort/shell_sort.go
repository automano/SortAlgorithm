package mysort

/*
1959年Shell发明，第一个突破O(n2)的排序算法，是简单插入排序的改进版。它与插入排序的不同之处在于，它会优先比较距离较远的元素。希尔排序又叫缩小增量排序。

算法描述
先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：

- 选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
- 按增量序列个数k，对序列进行k 趟排序；
- 每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

希尔排序的核心在于间隔序列的设定。既可以提前设定好间隔序列，也可以动态的定义间隔序列。动态定义间隔序列的算法是《算法（第4版）》的合著者Robert Sedgewick提出的。
*/

// ShellSort 希尔排序
// It is invented by Shell in 1959. It is an improved sorting algorithm on the base of the simple insertion sort.
// It differs from insertion sort in that it prioritizes the elements that are farther away. Shell sort is also called Diminishing Increment Sort.
func ShellSort(data []int) {
	for gap := len(data) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(data); i++ {
			j := i
			cur := data[i]
			for j-gap >= 0 && cur < data[j-gap] {
				data[j] = data[j-gap]
				j = j - gap
			}
			data[j] = cur
		}
	}
}
