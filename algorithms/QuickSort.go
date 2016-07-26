package main

import "fmt"


func quickSort(arr []int64, low int, high int) {
	if low >= high {
		return
	}
	middle := low + (high - low) / 2
	pivot := arr[middle]
	i, j := low, high
	for i < j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--
		}
	}
	if low < j {
		quickSort(arr, low, j)
	}
	if high > i {
		quickSort(arr, i, high)
	}
}

func output(arr []int64) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("\n%v	", arr[i])
	}
}

func main() {
	var arr = []int64{9, 2, 4, 7, 3, 7, 10};
	quickSort(arr, 0, len(arr) - 1)
	output(arr)
}
