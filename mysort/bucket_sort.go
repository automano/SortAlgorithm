package mysort

import "math"

// 9、桶排序（Bucket Sort）
// 桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。

// 9.1 算法描述
// - 设置一个定量的数组当作空桶；
// - 遍历输入数据，并且把数据一个一个放到对应的桶里去；
// - 对每个不是空的桶进行排序；
// - 从不是空的桶里把排好序的数据拼接起来。

// BucketSort - 桶排序
func BucketSort(data []int, bucketSize int) []int {
	// 确定数据范围
	max, min := math.MinInt32, math.MaxInt32
	for _, n := range data {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	// 初始化桶
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}
	// 按照规则放入对应的桶中
	for _, n := range data {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}
	//
	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			InsertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}
	return sorted
}
