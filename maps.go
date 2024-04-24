// package main

// import "fmt"

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func main(){
// 	m = make(map[string]Vertex)


// 	m["Bell Labs"] = Vertex{
// 		40.68433, -74.39967,
// 	}

// 	m["google"] = Vertex{29,90}  

// 	value, ok := m["Bell Labs"]
// if ok {
//     fmt.Println("Key 'Bell Labs' is present in the map. Value:", value)
// } else {
//     fmt.Println("Key 'Bell Labs' is not present in the map.")
// }

// }