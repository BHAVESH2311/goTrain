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
	for i:=0;i<s;i++{
		var num int
		fmt.Scan(&num)
		arr=append(arr,num)
	}
	intSlice:=sort.IntSlice(arr)
	intSlice.Sort()

	counts := make(map[int]int)

      for _, v := range arr {
          count, ok := counts[v] 
          if ok {
              counts[v] = count + 1  
          } else {    
              counts[v] = 1
          }    
      }

      for k, v := range counts {
          fmt.Printf("%d occurs %d times\n", k, v)  
      }    




}