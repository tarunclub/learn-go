package main

import "fmt"

// var i int = 78

func main(){
	var i int = 42
	j := 56
	fmt.Printf("The value of i is %v and type is %T\n",i,i)
	fmt.Printf("The value of j is %v and type is %T\n",j,j)

	var k float32
	k = float32(i) //coversion operator
	fmt.Println(k)
}
