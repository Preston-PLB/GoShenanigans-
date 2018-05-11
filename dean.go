package main

import ("fmt"
	"math/rand"
	"time")

func main() {
	start := [5]string{"%s 'em up", "mother %s", "son of a %s", "mother of %s", "%ser"}
	end := [4]string{"frick","dang","crap","father"}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i:=0; i<10; i++{
		fmt.Printf(start[r1.Intn(5)]+"\n", end[r1.Intn(3)])
	}
}