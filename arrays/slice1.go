package main 


import "fmt"



func main() {

	s := make([]string, 5)
	
	fmt.Println("emp: ", s)

	s[0] = "H"
	s[1] = "e" 
	s[2] = "l"
	s[3] = "l"
	s[4] = "o"


	fmt.Println("get: ", s)

	// appending to the list 
	s = append(s, " ")
	s = append(s, "w")
	s = append(s, "o")
	s = append(s, "r")
	s = append(s, "l")
	s = append(s, "d")

	fmt.Println("apd: ", s)

	// slicing 
	z := s[2:4]
	fmt.Println("slc: ", z)

	z = s[0:5]
	fmt.Println("slc: ", z)

}