package main

import "fmt"

type Vertex struct {
	X int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{22, 33}
	fmt.Println(v.X)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.y)
	p := &v
	fmt.Println(p.X)
	p.y = 44444
	fmt.Println(p.y)
}
