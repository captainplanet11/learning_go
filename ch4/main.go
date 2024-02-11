package main

// func main() {

// 	x := 10
// 	if x > 5 {
// 		fmt.Println(x)
// 		x := 5
// 		fmt.Println(x)
// 	}
// 	fmt.Println(x)
// }

// shadowinwf with multiple assigment
// func main() {
// 	x := 10
// 	if x > 5 {
// 		x, y := 5, 20
// 		fmt.Println(x, y)
// 	}
// 	fmt.Println(x)
// }

// shadowing true

// func main() {

// 	fmt.Println(true)
// 	true := 10
// 	fmt.Println(true)
// }

//if and else rule

// func main() {
// 	if n := rand.Intn(10); n == 0 {
// 		fmt.Println("That's too low")
// 	} else if n > 5 {
// 		fmt.Println("That's too big:", n)
// 	} else {
// 		fmt.Println("That's a good number:", n)
// 	}
// }

//Out of Scope…

// func main() {
// 	if n := rand.Intn(10); n == 0 {
// 		fmt.Println("That's too low")
// 	} else if n > 5 {
// 		fmt.Println("Thats too big: ", n)
// 	} else {
// 		fmt.Println("That's a good number : ", n)
// 	}
// 	fmt.Println(n)
// }

// example a complete for statement

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// // 	}
// // }

// func main() {
// 	evenVals := []int{2, 4, 6, 8, 10, 12}
// 	for i, v := range evenVals {
// 		fmt.Println(i, v)
// 	}
// }

// func main() {
// 	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
// 	for k := range uniqueNames {
// 		fmt.Println(k)
// 	}
// }

// func main() {
// 	m := map[string]int{
// 		"a": 1,
// 		"b": 2,
// 		"c": 3,
// 	}

// 	for i := 0; i < 3; i++ {
// 		fmt.Println("Loop", i)
// 		for k, v := range m {
// 			fmt.Println(k, v)
// 		}
// 	}
// }

// iterating over string
// func main() {
// 	samples := []string{"hello", "apple_π!"}
// 	for _, sample := range samples {
// 		for i, r := range sample {
// 			fmt.Println(i, r, string(r))
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	evenVals := []int{2, 4, 6, 8, 10, 12}
// 	for _, v := range evenVals {
// 		v *= 2
// 		fmt.Println(v)
// 	}
// 	fmt.Println(evenVals)
// }

// func main() {
// 	evenVals := []int{2, 4, 6, 8, 10}
// 	for i := 1; i < len(evenVals)-1; i++ {
// 		fmt.Println(i, evenVals[i])
// 	}

// }

// func main() {
// 	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropoloigst"}
// 	for _, word := range words {
// 		switch size := len(word); size {
// 		case 1, 2, 3, 4:
// 			fmt.Println(word, "is a short word")
// 		case 5:
// 			wordLen := len(word)
// 			fmt.Println(word, "is exactly the right length: ", wordLen)
// 		case 6, 7, 8, 9:
// 		default:
// 			fmt.Println(word, "is a long word")
// 		}
// 	}

// }

// go to yes gptp
