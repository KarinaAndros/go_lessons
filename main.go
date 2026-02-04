package main

//import packages
import "fmt"

//start program
func main(){
	//output text with new line
	fmt.Println("Hello, mzfk")
	
	//output text without  new line
	fmt.Print("Meow")

	//variables
	// score := 0
	// score++ //increment
	// score-- //decrement

	var number int
	var text string
	fmt.Println("number:", number)
	fmt.Println(text)
	// fmt.Println(score)

	//if else
	score := 5
	if score > 10{
		fmt.Println("Good")
	}else{
		fmt.Println("Bad")
	}
}
