package main

import (
	"fmt"
	"sort"
)

func main() {
	// Test data
	testData := []int{64, 34, 25, 12, 22, 11, 90}
	
	fmt.Println("Original array:", testData)
	
	// Test BubbleSort
	sorted := sort.BubbleSort(copySlice(testData))
	fmt.Println("BubbleSort result:", sorted)
	
	// Test SelectionSort
	sorted = sort.SelectionSort(copySlice(testData))
	fmt.Println("SelectionSort result:", sorted)
	
	// Test InsertionSort
	sorted = sort.InsertionSort(copySlice(testData))
	fmt.Println("InsertionSort result:", sorted)
	
	// Test MergeSort
	sorted = sort.MergeSort(copySlice(testData))
	fmt.Println("MergeSort result:", sorted)
	
	// Test QuickSort
	sorted = sort.QuickSort(copySlice(testData))
	fmt.Println("QuickSort result:", sorted)
	
	// Verify that original data is unchanged
	fmt.Println("Original array after tests:", testData)
}

func copySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}