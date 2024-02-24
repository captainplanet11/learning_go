package main

// import (
// 	"fmt"
// 	"./formatter"
// 	"./math"
// )

// func main() {
// 	num := math.Double(2)s
// 	output := print.Format(num)
// 	fmt.Println(output)
// }

// package main

// import (
// 	crand "crypto/rand"
// 	"encoding/binary"
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	r := seedRand()
// 	fmt.Println(r.Int())
// }

// func seedRand() *rand.Rand {
// 	var b [8]byte
// 	_, err := crand.Read(b[:])
// 	if err != nil {
// 		panic("cannot seed with cryptographic random number generator")
// 	}
// 	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
// 	return r
// }

// Many represents the combination of an amount of moeny
// and the currency the moneuy is in

// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns a
// the converted currency and any errors encountered from unknown or
// unconvertable
// currencies.
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
// USD - US Dollar
// CAD - Canadian Dollar
// EUR - Euro
// INR - Indian Rupee
//
// More information on exchange rates can be found
// at https://www.investopedia.com/terms/e/exchangerate.asp
// func Convert(from Money, to string) (Money, error) {
// // ...
// }
