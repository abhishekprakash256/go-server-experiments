package main 

import "fmt"
import "unsafe"


func add(a int , b int) (int) {
	
	var sum int = a + b 
	
	return sum
}



func main() {
	// make variable

	var a int = 10
	var b int = 20

	city := "New York"

	// print 

	fmt.Println("integer a is: ", a)	
	fmt.Println("Integer sum is : ", a+b)
	fmt.Println("City is: ", city)
	// print type of variable
	fmt.Printf("Type of a is: %T\n", a)
	fmt.Printf("Type of b is: %T\n", b)
	fmt.Printf("Type of city is: %T\n", city)
	
	// print size of variable
	fmt.Printf("Size of a is: %d\n", unsafe.Sizeof(a))
	fmt.Printf("Size of b is: %d\n", unsafe.Sizeof(b))
	fmt.Printf("Size of city is: %d\n", unsafe.Sizeof(city))


	res := add(a,b) 
	// get the sum
	fmt.Println("Sum by the func ", res )



}