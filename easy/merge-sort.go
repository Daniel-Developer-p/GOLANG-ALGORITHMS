package easy

import (
	"fmt"
)

func merge(a []int, b []int) []int {
	var r = make([]int, len(a) + len(b))
	var i, j = 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i + j] = a[i]
			i++
		} else {
			r[i + j] = b[j]
			j++
		}
	}

	for i < len(a) { r[i + j] = a[i]; i++ }
	for j < len(a) { r[i + j] = b[j]; j++ }

	return r
}

func MergeSort(items []int) []int {
	if len(itmes) < 2 {
		return items
	}

	var middle = len(items) / 2
	var a = MergeSort(items[:middle])
	var b = MergeSort(items[middle:])
	
	return merge(a, b)
}

func main() {
	fmt.Println(MergeSort([]int { 10, 9, 8, 7, 6, 5, 4, 3, 2, 1 }), )
}