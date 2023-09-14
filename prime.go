package main

import "fmt"

func main(){

	fmt.Println("Enter a number")
	var num int;
	count:=0;
	fmt.Scan(&num)
	if(num==2){
		fmt.Println("The number is a prime number")
	}
	for i:=2;i<num;i++ {
		if(num%i==0){
			count++
		}
	}

	if(count==0){
		fmt.Println("The number is a prime number")
	}else{
		fmt.Println("the number is not a prime number")
	}
}