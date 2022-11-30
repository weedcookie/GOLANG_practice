package main 

import "fmt"


func main() {

	var x [2][2]int
	for i :=0; i< 2; i++ {
		for j :=0; j < 2; j++ {
			x[i][j] = i+j
		}
	}

	fmt.Println("2d: ", x)

}