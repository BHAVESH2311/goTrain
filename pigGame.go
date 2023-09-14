package main
import (
	"fmt"
	"math/rand"
)
func main(){
	fmt.Println("Welcome to pig game")
	var totalScore int =0
	var turnCount int =1
	var turnScore int = 0
	for(totalScore<=20 && turnScore<=20){
	fmt.Println("press h to hold and r to roll")
		var choice string 
		fmt.Scan(&choice)
		switch choice {
		case "r":
			randomNumber := rand.Intn(6) + 1
			fmt.Println("your score is ",randomNumber)
			if(randomNumber==1){
				turnScore =0
				fmt.Println(turnScore)
				turnCount++	
			}else{
				turnScore+=randomNumber
				fmt.Println(turnScore)	
			}
		case "h":
			turnCount++
			totalScore+=turnScore
			fmt.Println("your totale score is ",totalScore)
			if(totalScore>=20){
				break
			}
		}
	}
	fmt.Println("Congrats you won in",turnCount)
	

}