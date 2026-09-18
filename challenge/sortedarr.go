package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4, 5}
	fmt.Printf("findMedianSortedArrays(nums1, nums2): %.2f\n", findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arrChan := make(chan []int)
	go func() {
		arr := append(nums1, nums2...)
		sort.Ints(arr)
		arrChan <- arr
	}()
	median := 0.0
	newArr := <- arrChan
	if len(newArr)%2 == 0 {
		half := len(newArr) / 2
		sum := newArr[half-1] + newArr[half]
		median = float64(sum) / 2.0
	} else {
		median = float64(newArr[len(newArr)/2])
	}
	return median
}
