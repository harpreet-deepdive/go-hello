package main

import (
	"fmt"
	"math"
	"strconv"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Area() float64{
return v.X * v.Y
}

func main() {

	c := Vertex{3,4}
	c.Abs()
	
	i, err := strconv.Atoi("42b")

	if(err != nil){
		fmt.Println("couldn't convert:", err)
	}else{

	fmt.Println(i)
}


	const age =  30



}
