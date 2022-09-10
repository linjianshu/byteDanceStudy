package main

import (
	"sort"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int{5, 2, 4, 6, 1, 3})
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort([]int{5, 2, 4, 6, 1, 3})
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort([]int{5, 2, 4, 6, 1, 3})
	}
}

func BenchmarkBobSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BobSort([]int{5, 2, 4, 6, 1, 3})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints([]int{5, 2, 4, 6, 1, 3})
	}
}
