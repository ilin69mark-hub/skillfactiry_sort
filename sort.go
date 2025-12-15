package sort

// BubbleSort sorts an array using bubble sort algorithm
func BubbleSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	
	result := make([]int, n)
	copy(result, arr)
	
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	
	return result
}

// SelectionSort sorts an array using selection sort algorithm
func SelectionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	
	result := make([]int, n)
	copy(result, arr)
	
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			result[i], result[minIdx] = result[minIdx], result[i]
		}
	}
	
	return result
}

// InsertionSort sorts an array using insertion sort algorithm
func InsertionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	
	result := make([]int, n)
	copy(result, arr)
	
	for i := 1; i < n; i++ {
		key := result[i]
		j := i - 1
		
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	
	return result
}

// MergeSort sorts an array using merge sort algorithm
func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	
	result := make([]int, n)
	copy(result, arr)
	
	mergeSortHelper(result, 0, n-1)
	
	return result
}

func mergeSortHelper(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		
		mergeSortHelper(arr, left, mid)
		mergeSortHelper(arr, mid+1, right)
		
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid
	
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)
	
	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}
	
	i, j, k := 0, 0, left
	
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// QuickSort sorts an array using quick sort algorithm
func QuickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	
	result := make([]int, n)
	copy(result, arr)
	
	quickSortHelper(result, 0, n-1)
	
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}