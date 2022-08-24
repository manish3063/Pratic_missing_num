package main

import "fmt"

func main() {
	arr := BubbleShort()

	result := missing(arr)
	fmt.Println(result)

}

func BubbleShort() []int {

	arr := []int{3, 0, 1}

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr)-1; j++ {

			if arr[j] > arr[j+1] {
				arr = Swap(arr, j, j+1)

			}

		}

	}
	fmt.Println(arr)
	return arr
}

func Swap(arr []int, a, b int) []int {

	temp := arr[a]

	arr[a] = arr[b]

	arr[b] = temp

	return arr

}
func missing(arr []int) int {
	for i := 0; i < len(arr); i++ {

		if i != arr[i] {

			return i
		}
	}
	return -1
}
