package heapsort

import (
	"reflect"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{12, 11, 13, 5, 6, 7}
	expected := []int{5, 6, 7, 11, 12, 13}
	heapsort(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestHeapSortItem(t *testing.T) {
	arr := []Item{
		{"a", 12},
		{"b", 11},
		{"c", 13},
		{"d", 5},
		{"e", 6},
		{"f", 7},
	}
	expected := []Item{
		{"d", 5},
		{"e", 6},
		{"f", 7},
		{"b", 11},
		{"a", 12},
		{"c", 13},
	}
	heapsortItem(arr)
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{12, 11, 13, 5, 6, 7}
		heapsort(arr)
	}
}

func BenchmarkSortInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{12, 11, 13, 5, 6, 7}
		sort.Ints(arr)
	}
}

func BenchmarkHeapSortItem(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []Item{
			{"a", 12},
			{"b", 11},
			{"c", 13},
			{"d", 5},
			{"e", 6},
			{"f", 7},
		}
		heapsortItem(arr)
	}
}

func BenchmarkSortSliceStable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []Item{
			{"a", 12},
			{"b", 11},
			{"c", 13},
			{"d", 5},
			{"e", 6},
			{"f", 7},
		}
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i].Value > arr[j].Value
		})
	}
}
