package easy

import (
	f "fmt"
)

func main(){
	arr := [9]int{2, 3, 5, 7, 9, 1, 4, 6, 8, 10}

	for out := 1; out < len(arr); out++ {
		temp := arr[out]
		in := out

		for ; in > 0 && arr[in-1] >= temp; in-- {
			arr[in] = arr[in-1]
		}
		arr[in] = temp
	}

	for _, sortedvals := range arr {
		f.Println(sortedvals)
	}
}