package main

import "fmt"

func binary_search(array []int, num int) bool {
	l, r := 0, len(array)
	answer := false
	var m int
	for {
		if l >= r {
			return answer
		}
		m = (l + r) / 2
		if num > array[m] {
			l = m + 1
		} else if num < array[m] {
			r = m
		} else {
			answer = true
			return answer
		}
	}
}

func create_array(start int, stop int) []int {
	var arr []int

	for i := start; i < stop+1; i++ {
		arr = append(arr, i)
	}

	return arr
}

func main() {
	arr := create_array(1, 100)

	fmt.Println("Введите число: ")
	var num int
	fmt.Scanf("%d\n", &num)

	num_in_array := binary_search(arr, num)

	fmt.Println(num_in_array)
}
