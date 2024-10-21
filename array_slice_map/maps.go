package main 

import "fmt"

func main(){


	// creating a map using a map literal 

	capitals := map[string]string{
		"France" :"Paris",
		"italy"  : "Rome",
		"Japan"  : "Tokyo",
	}


	fmt.Println("capitals", capitals)

	// creating map using make 

	scores := make(map[string]int)
	scores["Alice"] = 90
	scores["Bob"] = 85
	scores["Charlie"] = 95
	fmt.Println("scores:", scores)

	// Accessing the value 

	aliceScore := scores["Alice"]
	fmt.Println("Alice's score :", aliceScore)



	// checking if  key exists 

	score, exists := scores["David"]
	if exists{
		fmt.Println("David's score:", score)
	}else{
		fmt.Println("david's score not found")
	}



	//Deleting a key-value pair 

	delete(scores,"Bob")

	fmt.Println("After deletion:", scores)

	//Iterating over a map 

	for country, capital := range capitals{
		fmt.Printf("The capital of %s is %s. \n", country, capital )
	}

}