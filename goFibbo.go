package main
import "fmt"


func main(){


	fmt.Println("Enter the length of fibbo series")

	var num int
	fmt.Scan(&num)

	t1:=0;
	t2:=1;
	sum:=0

	for i:=0;i<num;i++{

		sum+=t1
		temp:=t1
		t1=t2
		t2=temp+t1
	}
	fmt.Println("The sum of",num, "fibbonacci numbers are ",sum)

}