package main 


import "fmt"

func main() {

	if condition := true; condition {
		fmt.Println("you can declare variables in if statemnts")
	} else {
		fmt.Println("Won't get executed ")
	}
}