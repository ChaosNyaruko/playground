package main

import "fmt"

func sumSubarrayMins(arr []int) (rst int) {
	n := len(arr)
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		for l[i] = i - 1; l[i] >= 0 && arr[l[i]] > arr[i]; l[i] = l[l[i]] {
		}
	}
	for i := n - 1; i >= 0; i-- {
		for r[i] = i + 1; r[i] < n && arr[r[i]] >= arr[i]; r[i] = r[r[i]] {
		}
	}
	for i, num := range arr {
		rst += num * (i - l[i]) * (r[i] - i)
		rst %= 1000000007
	}
	return
}
func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
}
