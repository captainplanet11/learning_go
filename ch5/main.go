package main

import (
	"fmt"
)

// func div(numerator int, denominator int) int {
// 	if denominator == 0 {
// 		return 0
// 	}
// 	return numerator / denominator
// }

// func main() {
// 	result := div(3, 2)
// 	fmt.Println(result)
// }

// type MyFuncOpts struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func MyFunc(opts MyFuncOpts) error {
// 	//do something
// }

// func main() {
// 	MyFunc(MyFuncOpts{
// 		LastName: "patel",
// 		Age:      50,
// 	})
// 	MyFunc(MyFuncOpts{
// 		FirstName: "Joe",
// 		LastName:  "Mama",
// 	})
// }

// func addTo(base int, vals ...int) []int {
// 	out := make([]int, 0, len(vals))
// 	for _, v := range vals {
// 		out = append(out, base+v)
// 	}
// 	return out
// }

// func main() {
// 	fmt.Println(addTo(3))
// 	fmt.Println(addTo(3, 2))
// 	fmt.Println(addTo(3, 2, 4, 6, 8))
// 	a := []int{4, 3}
// 	fmt.Println(addTo(3, a...))
// 	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
// }

// func divAndRemainder(numerator int, denominator int) (int, int, error) {
// 	if denominator == 0 {
// 		return 0, 0, errors.New("cannot divide by zero")
// 	}
// 	return numerator / denominator, numerator % denominator, nil
// }

// func main() {
// 	result, remainder, err := divAndRemainder(5, 2)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(result, remainder)
// }

// func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
// 	if denominator == 0 {
// 		err = errors.New("cannot divide by zero")
// 		return result, remainder, err
// 	}
// 	result, remainder = numerator/denominator, numerator%denominator
// 	return result, remainder, err
// }

// func main() {
// 	res, rem, err := divAndRemainder(5, 2)
// 	fmt.Println(res, rem, err)
// }

//functions are values
// primitive calculator

// func add(i int, j int) int { return i + j }

// func sub(i int, j int) int { return i - j }

// func mul(i int, j int) int { return i * j }

// func div(i int, j int) int { return i / j }

// var opMap = map[string]func(int, int) int{
// 	"+": add,
// 	"-": sub,
// 	"*": mul,
// 	"/": div,
// }

// func main() {
// 	expressions := [][]string{
// 		[]string{"2", "+", "3"},
// 		[]string{"2", "-", "3"},
// 		[]string{"2", "*", "3"},
// 		[]string{"2", "/", "3"},
// 		[]string{"2", "%", "3"},
// 		[]string{"two", "+", "three"},
// 		[]string{"5"},
// 	}
// 	for _, expression := range expressions {
// 		if len(expression) != 3 {
// 			fmt.Println("invalid expression:", expression)
// 			continue
// 		}
// 		p1, err := strconv.Atoi(expression[0])
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		op := expression[1]
// 		opFunc, ok := opMap[op]
// 		if !ok {
// 			fmt.Println("unsupported operator:", op)
// 			continue
// 		}
// 		p2, err := strconv.Atoi(expression[2])
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		result := opFunc(p1, p2)
// 		fmt.Println(result)
// 	}
// }

// Anonymous Functions

// func main() {
// 	for i := 0; i < 5; i++ {
// 		func(j int) {
// 			fmt.Println("Printing", j, "from inside the anonymous function")
// 		}(i)
// 	}
// }

// func main() {
// 	type Person struct {
// 		FirstName string
// 		LastName  string
// 		Age       int
// 	}

// 	people := []Person{
// 		{"Pat", "Patterson", 37},
// 		{"Tracy", "Bobbert", 23},
// 		{"Fred", "Fredson", 18},
// 	}
// 	fmt.Println(people)

// 	// sort by last name
// 	sort.Slice(people, func(i int, j int) bool {
// 		return people[i].LastName < people[j].LastName
// 	})
// 	fmt.Println(people)

// 	// sort by age
// 	sort.Slice(people, func(i int, j int) bool {
// 		return people[i].Age < people[j].Age
// 	})
// 	fmt.Println(people)
// }

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}
