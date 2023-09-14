package main
import "fmt"
import "time"

func main() {

	now:=time.Now()

	if now.Hour() >= 6 && now.Hour() < 11 {
		fmt.Println("Good morning!")
	} else if now.Hour() >= 11 && now.Hour() < 16 {
		fmt.Println("Good afternoon!")        
	} else if now.Hour() >= 16 && now.Hour() < 21 {
		fmt.Println("Good evening!")
	} else {
		fmt.Println("Good night!")
	}

}