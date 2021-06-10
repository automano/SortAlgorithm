package mysort

/*
堆排序（Heap Sort）
堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。

算法描述
- 将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
- 将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
- 由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。
- 不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
*/

// HeapSort - 堆排序
func HeapSort(data []int) {
	// 建堆 - 无序数据
	heapify(data)
	for i := len(data) - 1; i > 0; i-- {
		// 将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
		data[0], data[i] = data[i], data[0]
		// 由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆
		siftDown(data, 0, i)
	}
}

func heapify(data []int) {
	for i := (len(data) - 1) / 2; i >= 0; i-- {
		siftDown(data, i, len(data))
	}
}

func siftDown(heap []int, lo, hi int) {
	root := lo
	for {
		child := root*2 + 1
		if child >= hi {
			break
		}
		if child+1 < hi && heap[child] < heap[child+1] {
			child++
		}
		if heap[root] < heap[child] {
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		} else {
			break
		}
	}
}
