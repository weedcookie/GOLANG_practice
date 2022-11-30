package main 

import "fmt"

func main () {
	fmt.Println("First loop ")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i +=1
	}	

	fmt.Println("Second loop")
	j := 2
	for {
		if j >= 4 {
			break
		}
		fmt.Println(j);
		j += 1 
	}
}