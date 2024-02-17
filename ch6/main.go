package main

// func main() {

// 	x := 10
// 	pointerToX := &x
// 	fmt.Println(pointerToX)  // prints a memory address
// 	fmt.Println(*pointerToX) // prints 10
// 	z := 5 + *pointerToX
// 	fmt.Println(z)

// }

// func main() {
// 	var x *int
// 	fmt.Println(x == nil) // prints true
// 	fmt.Println(*x)       // panics

// }

// func failedUpdate(g *int) {
// 	x := 10
// 	g = &x
// }

// func main() {
// 	var f *int // f is nil
// 	failedUpdate(f)
// 	fmt.Println(f) //prints nil

// }

// func failedUpdate(px *int) {
// 	x2 := 20
// 	px = &x2

// }

// func update(px *int) {
// 	*px = 20
// }

// func main() {
// 	x := 10
// 	failedUpdate(&x)
// 	fmt.Println(x) //prints 10
// 	update(&x)
// 	fmt.Println(x) //prints 20
// }
