package main 
import "fmt"


func main() {
	sum :=0


	// declaring for loops 
	for i := 1; i < 5; i++ {
		sum += i
	}

	fmt.Println(sum)
}