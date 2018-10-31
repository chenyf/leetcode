package main

import "fmt"

func quicksort(l []int) []int {
	if l == nil || len(l) == 1 {
		return l
	}
	_quicksort(l, 0, len(l)-1)
	return l
}

func _quicksort(l []int, start, end int) {
	if start >= end {
		return
	}

	// 先分区，再针对两个分区递归排序
	q := _partition(l, start, end)
	_quicksort(l, start, q-1)
	_quicksort(l, q+1, end)
}

func _partition(l []int, start, end int) int {
	pivot := l[end]
	i := start
	for j := start; j < end; j++ {
		if l[j] < pivot {
			if j != i {
				t := l[i]
				l[i] = l[j]
				l[j] = t
			}
			i += 1
		}
	}

	t := l[i]
	l[i] = l[end]
	l[end] = t
	return i
}

func mergesort(l []int) []int {
	if l == nil || len(l) == 1 {
		return l
	}
	_mergesort(l, 0, len(l)-1)
	return l
}

func _mergesort(l []int, start, end int) {
	if start >= end {
		return
	}

	q := (start + end) / 2
	_mergesort(l, start, q)
	_mergesort(l, q+1, end)
	_merge(l, start, end, q)
}

func _merge(l []int, start, end, q int) {

	fmt.Printf("before merge: %d %d %d\n", start, end, q)
	fmt.Println(l)

	var result = make([]int, len(l))

	i, j, k := start, q+1, 0

	for i <= q && j <= end {
		if l[i] <= l[j] {
			result[k] = l[i]
			i += 1
		} else {
			result[k] = l[j]
			j += 1
		}
		k += 1
	}

	m := i
	n := q
	if j <= end {
		m = j
		n = end
	}

	for m <= n {
		result[k] = l[m]
		k += 1
		m += 1
	}

	for i := 0; i < k; i++ {
		l[start+i] = result[i]
	}
	fmt.Printf("after merge\n")
	fmt.Println(l)
}

func main() {
	l := []int{7, 8, 9, 10, 0, 2, 5, 4, 2, 1, 3}
	fmt.Println(mergesort(l))
	//fmt.Println(quicksort(l))
}
