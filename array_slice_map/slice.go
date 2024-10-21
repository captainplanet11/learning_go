package main 

import "fmt"


func main(){

	languages := []string{"Go", "Python", "Java"}
    fmt.Println("Languages:", languages)

    fmt.Println("First language:", languages[0])

    //slicing a slice 
    subSlice := languages[1:3]
    fmt.Println("sub-slice", subSlice)

    //using make to create a slcie

    numbers := make([] int,3,5) //len 3, cap = 5
    numbers[0] = 100
    numbers[1] = 200
    numbers[2] = 300

    fmt.Println("Nmbers:", numbers)
    fmt.Println("Lenght:", len(numbers))
    fmt.Println("Capacity:", cap(numbers))

   //appending on a slice

   numbers  =  append(numbers,400,500)
   fmt.Println("After append", numbers)

   //appending beyond the capacity 

   numbers = append(numbers,600)
   fmt.Println("After exceeding capacity", numbers)
   fmt.Println("New capacity", cap(numbers))



}