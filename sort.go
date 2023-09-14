package main

import "fmt"
import "sort"

func main() {

	fmt.Println("Enter the array size")

	var s int
	fmt.Scan(&s)

	var arr []int

	for i := 0; i < s; i++ {
		var num int  
		fmt.Scan(&num)  
		arr = append(arr, num)
	}

	intSlice := sort.IntSlice(arr)
	intSlice.Sort()

	fmt.Println("The sorted array is:")
	for i := 0; i < s; i++ {
		fmt.Println(arr[i])
	}

}