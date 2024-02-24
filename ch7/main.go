package main

//structs
// pointers and value recievers

// type Counter struct {
// 	total       int
// 	lastUpdated time.Time
// }

// func (c *Counter) Increment() {
// 	c.total++
// 	c.lastUpdated = time.Now()
// }

// func (c Counter) String() string {
// 	return fmt.Sprintf("total: %d, last updated: %v", c.total,
// 		c.lastUpdated)
// }

// func main() {
// 	var c Counter
// 	fmt.Println(c.String())
// 	c.Increment()
// 	fmt.Println(c.String())
// }

// func doUpdateWrong(c Counter) {
// 	c.Increment()
// 	fmt.Println("in doUpdateWrong:", c.String())
// 	}
// 	func doUpdateRight(c *Counter) {
// 	c.Increment()
// 	fmt.Println("in doUpdateRight:", c.String())
// 	}
// 	func main() {
// 	var c Counter
// 	doUpdateWrong(c)
// 	fmt.Println("in main:", c.String())
// 	doUpdateRight(&c)
// 	fmt.Println("in main:", c.String())
// 	}

// passing on a nil value

// type IntTree struct {
// 	val         int
// 	left, right *IntTree
// }

// func (it *IntTree) Insert(val int) *IntTree {
// 	if it == nil {
// 		return &IntTree{val: val}
// 	}
// 	if val < it.val {
// 		it.left = it.left.Insert(val)
// 	} else if val > it.val {
// 		it.right = it.right.Insert(val)
// 	}
// 	return it
// }

// func (it *IntTree) Contains(val int) bool {
// 	switch {
// 	case it == nil:
// 		return false
// 	case val < it.val:
// 		return it.left.Contains(val)
// 	case val > it.val:
// 		return it.right.Contains(val)
// 	default:
// 		return true
// 	}
// }

// func main() {
// 	var it *IntTree
// 	it = it.Insert(5)
// 	it = it.Insert(3)
// 	it = it.Insert(10)
// 	it = it.Insert(2)
// 	fmt.Println(it.Contains(2))
// 	fmt.Println(it.Contains(12))
// }

// type Adder struct {
// 	start int
// }

// func (a Adder) AddTo(val int) int {
// 	return a.start + val
// }

// func main() {
// 	myAdder := Adder{start: 10}
// 	fmt.Println(myAdder.AddTo(5))
// 	f1 := myAdder.AddTo
// 	fmt.Println(f1(10))
// 	f2 := Adder.AddTo
// 	fmt.Println(f2(myAdder, 15))
// }

// assigning untyped constants is valid

// iota

// const (
// 	Sunday    = iota // 0
// 	Monday           // 1
// 	Tuesday          // 2
// 	Wednesday        // 3
// 	Thursday         // 4
// 	Friday           // 5
// 	Saturday         // 6
// )

// func main() {
// 	fmt.Println("Every day of the week: ")
// 	fmt.Println("Sunday:", Sunday)
// 	fmt.Println("Monday:", Monday)
// 	fmt.Println("Tuesday:", Tuesday)
// 	fmt.Println("Wednesday:", Wednesday)
// 	fmt.Println("Thursday:", Thursday)
// 	fmt.Println("Friday:", Friday)
// 	fmt.Println("Saturday:", Saturday)
// }

// func main() {
// 	const (
// 		KB = 1 << (10 * iota) // 1 << (10*0) = 1
// 		MB                    // 1 << (10*1) = 1024
// 		GB                    // 1 << (10*2) = 1048576
// 		TB                    // 1 << (10*3) = 1073741824
// 	)
// 	fmt.Println("KB:", KB)
// 	fmt.Println("MB:", MB)
// 	fmt.Println("GB:", GB)
// 	fmt.Println("TB:", TB)
// }

// Shape interface defines a method for calculating area
// type Shape interface {
//     Area() float64
// }

// // Circle struct represents a circle
// type Circle struct {
//     Radius float64
// }

// // Area calculates the area of the circle
// func (c Circle) Area() float64 {
//     return math.Pi * c.Radius * c.Radius
// }

// // Rectangle struct represents a rectangle
// type Rectangle struct {
//     Width, Height float64
// }

// // Area calculates the area of the rectangle
// func (r Rectangle) Area() float64 {
//     return r.Width * r.Height
// }

// func main() {
//     // Both Circle and Rectangle implement the Shape interface
//     circle := Circle{Radius: 5}
//     rectangle := Rectangle{Width: 3, Height: 4}

//     // Calculate and print the areas
//     fmt.Println("Circle area:", circle.Area())
//     fmt.Println("Rectangle area:", rectangle.Area())
// }
