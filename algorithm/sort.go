package main

import "fmt"

func main() {
	sortArr := []int{6, 4, 12, 3, 1, 5, 8}
	//SortChoose(sortArr)
	SortInsert(sortArr)
	for i, v := range sortArr {
		fmt.Println(i, v)
	}
}

// SortChoose sort method of choose
func SortChoose(unsort []int) {
	for i := 0; i < len(unsort); i++ {
		minidx := i
		for j := i + 1; j < len(unsort); j++ {
			if unsort[j] < unsort[minidx] {
				minidx = j
			}
		}
		if minidx != i {
			unsort[i], unsort[minidx] = unsort[minidx], unsort[i]
		}
	}
}

func SortInsert(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list)-i; j++ {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
}
