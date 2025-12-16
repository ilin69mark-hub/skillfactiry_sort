package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSortingAlgorithms(t *testing.T) {
	testData := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}

	// Test BubbleSort
	result := BubbleSort(testData)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSort failed. Expected %v, got %v", expected, result)
	}

	// Test SelectionSort
	result = SelectionSort(testData)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SelectionSort failed. Expected %v, got %v", expected, result)
	}

	// Test InsertionSort
	result = InsertionSort(testData)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InsertionSort failed. Expected %v, got %v", expected, result)
	}

	// Test MergeSort
	result = MergeSort(testData)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeSort failed. Expected %v, got %v", expected, result)
	}

	// Test QuickSort
	result = QuickSort(testData)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("QuickSort failed. Expected %v, got %v", expected, result)
	}

	// Verify original slice is unchanged
	if !reflect.DeepEqual(testData, []int{64, 34, 25, 12, 22, 11, 90}) {
		t.Error("Original slice was modified")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			BubbleSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkSelectionSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			InsertionSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkMergeSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkQuickSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			QuickSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
