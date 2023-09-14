package main

import "fmt"

func main(){

	var eve =0
	var odd =0
	var zero=0

	fmt.Println("enter the size of arr")

	var s int

	
	fmt.Scan(&s)


	fmt.Println("please enter the elements of arrays")

	var arr []int

	for i:=0;i<s;i++{
		var num int
		fmt.Scan(&num)
		arr=append(arr,num)
	}

	for i:=0;i<s;i++{

		if(arr[i]==0){
			zero++
			continue
		}
		if(arr[i]%2==0 && arr[i]!=0){
			eve++
		}else{
			odd++
		}
	}


	
	fmt.Println("The count of even number in your given array is",eve)
	fmt.Println("The count of odd number in your given array is",odd)
	fmt.Println("The count of zeroes in your given array is",zero)


}