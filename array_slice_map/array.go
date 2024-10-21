package main 

import "fmt"


func main(){
	// Declare an araray of 5 integers 
	var  numbers [5]int 
	numbers[0] = 10
	numbers[1] = 20
    numbers[2] = 30
    numbers[3] = 40
    numbers[4] = 50

    fmt.Println("Array:",numbers)

    //initialize an array with values

    fruits := [5]string{"Apple","Banana","pineappple"}
    fruits[3] = "mango"
    fmt.Println("fruits", fruits)

}