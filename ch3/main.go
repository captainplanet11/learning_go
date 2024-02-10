package main

import "fmt"

// func main() {
// 	var x = [...]int{1, 2, 3, 4, 5}
// 	var y = [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(x == y)
// }

// func main() {
// 	var x = []int{1, 2, 3}
// 	x = append(x, 4)
// 	fmt.Println(x)
// }

// func main() {
// 	var x []int
// 	fmt.Println(x, len(x), cap(x))
// 	x = append(x, 10)
// 	fmt.Println(x, len(x), cap(x))
// 	x = append(x, 20)
// 	fmt.Println(x, len(x), cap(x))
// 	x = append(x, 30)
// 	fmt.Println(x, len(x), cap(x))
// 	x = append(x, 40)
// 	fmt.Println(x, len(x), cap(x))
// 	x = append(x, 50)
// 	fmt.Println(x, len(x), cap(x))
// }

// func main() {

// 	x := []int{1, 2, 3, 4}
// 	y := x[:2]
// 	z := x[1:]
// 	d := x[1:3]
// 	e := x[:]
// 	fmt.Println("x:", x)
// 	fmt.Println("y:", y)
// 	fmt.Println("z:", z)
// 	fmt.Println("d:", d)
// 	fmt.Println("e:", e)

// }

// func main() {

// 	x := [4]int{5, 6, 7, 8}
// 	y := x[:2]
// 	z := x[2:]
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }

// func main() {
// 	x := []int{1, 2, 3, 4}
// 	y := make([]int, 4)
// 	num := copy(y, x)
// 	fmt.Println(y, num)
// }

// func main() {
// 	var s string = "Hello there"
// 	var s2 string = s[4:7]
// 	var s3 string = s[:5]
// 	var s4 string = s[6:]
// 	fmt.Println(s2, s3, "", s4)
// }

// func main (){
// var s string = "Hello 🌞"
// var s2 string = s[4:7]
// var s3 string = s[:5]
// var s4 string = s[6:]
// }

// func main() {
// 	var s string = "Hello, 🌞"
// 	var bs []byte = []byte(s)
// 	var rs []rune = []rune(s)
// 	fmt.Println(bs)
// 	fmt.Println(rs)
// }

// func main() {
// var x int = 65
// var y = string(x)
// fmt.Println(y)
// }

// func main() {

// 	//map[KeyType]valueType
// 	//var nilMap map[string]int
// 	//totalWins := map[string]int{}

// 	teams := map[string][]string{
// 		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
// 		"Lions":   []string{"Sarah", "Peter", "Billie"},
// 		"Kittens": []string{"Waldo", "Raul", "Ze"},
// 	}

// 	fmt.Println(teams)

// }

// func main() {
// 	totalWins := map[string]int{}
// 	totalWins["Orcas"] = 1
// 	totalWins["Lions"] = 2
// 	fmt.Println(totalWins["Orcas"])
// 	fmt.Println(totalWins["Kittens"])
// 	totalWins["Kittens"]++
// 	fmt.Println(totalWins["Kittens"])
// 	totalWins["Lions"] = 3
// 	fmt.Println(totalWins["Lions"])
// }

// func main() {
// 	m := map[string]int{
// 		"hello": 5,
// 		"world": 10,
// 	}
// 	fmt.Println(m["hello"])
// 	delete(m, "hello")
// 	fmt.Println(m)

// }

// func main() {
// 	intSet := map[int]bool{}
// 	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
// 	for _, v := range vals {
// 		intSet[v] = true
// 	}
// 	fmt.Println(len(vals), len(intSet))
// 	fmt.Println(intSet[5])
// 	fmt.Println(intSet[500])
// 	if intSet[100] {
// 		fmt.Println("100 is in the set")
// 	}
// }

// structs ----------------------------------------------------------------------------

// func main() {
// 	type person struct {
// 		name string
// 		age  int
// 		pet  string
// 	}

// 	bob := person{}

// 	julia := person{
// 		"Julia",
// 		40,
// 		"cat",
// 	}

// 	fmt.Println(bob)
// 	fmt.Println(julia)
// 	bob.name = "Bob"
// 	fmt.Println(bob)
// }

func main() {
	type firstPerson struct {
		name string
		age  int
	}

	f := firstPerson{
		name: "BOb",
		age:  50,
	}

	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g)

}
