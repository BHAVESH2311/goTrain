package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("enter the size of arr")
	var s int
	fmt.Scan(&s)
	fmt.Println("please enter the elements of arrays")
	var arr []int
	var count=0
	for i:=0;i<s;i++{
		var num int
		fmt.Scan(&num)
		arr=append(arr,num)
	}
	intSlice := sort.IntSlice(arr)
	sort.Sort(sort.Reverse(intSlice))
	fmt.Println("The sorted array is:")
	for i := 0; i < s; i++ {
		fmt.Println(arr[i])
	}
	for i:=0;i<s;i++{
		if(arr[i]==arr[i+1]){
			count++
		}else{
			break
		}
	}
	if(count!=0){	
	fmt.Println(count)
	fmt.Println("Second Largest Number is",arr[count+1])
	}else{
		fmt.Println("second largest number is",arr[1])
	}
}