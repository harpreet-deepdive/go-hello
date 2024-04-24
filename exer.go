// package main

// import (
// 	"fmt"
// 	"strings"
// )



// func WordCount(s string) map[string]int {

// 	var stringArr = strings.Fields(s)
	
// 	  var m = make(map[string]int)

// 		for i := 0; i < len(stringArr); i++ {
// 			m[stringArr[i]]++
// 		}


// 	return m
// }
// func incrementer() func() int {
//     a,b:=0,1
//     return func() int {
// 		result := a
// 		a, b = b, a+b
// 		return result
//     }
// }

// func main() {
//     f := incrementer()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }


// // func main() {
// //  fmt.Println(WordCount("Go I am Learning"))	
// // }
