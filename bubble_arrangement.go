// сортировка пузырьком на языке Go
package main

import "fmt"

func main() {
	arr := [5]int{5, 4, 3, 2, 1}
	for i := 0; i < (len(arr) - 1); i++ {
		for j := 0; j < (len(arr) - 1 - i); j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println(arr)
		}
	}
	fmt.Println(arr)
}
