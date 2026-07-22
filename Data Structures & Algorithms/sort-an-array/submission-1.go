func sortArray(arr []int) []int {
    arr = maxHeapify(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapifyDown(arr[:i], 0)
	}
	return arr
}

func leftChildIndex(index int) int {
	return 2*index + 1
}

func rightChildIndex(index int) int {
	return 2*index + 2
}

func lastNonLeaf(arr []int) int {
	return len(arr)/2 - 1
}


func heapifyDown(arr []int, index int) []int {
	for {
		leftIndex := leftChildIndex(index)
		rightIndex := rightChildIndex(index)
		largestIndex := index

		if leftIndex < len(arr) && arr[leftIndex] > arr[largestIndex] {
			largestIndex = leftIndex
		}
		if rightIndex < len(arr) && arr[rightIndex] > arr[largestIndex] {
			largestIndex = rightIndex
		}
		if largestIndex != index {
			arr[index], arr[largestIndex] = arr[largestIndex], arr[index]
			index = largestIndex
		} else {
			break
		}

	}
	return arr
}

func maxHeapify(arr []int) []int{
	nonLeaf := lastNonLeaf(arr)
	for i := nonLeaf; i >= 0; i-- {
		arr = heapifyDown(arr, i)
	}
	return arr
}
