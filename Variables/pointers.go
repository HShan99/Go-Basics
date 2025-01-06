package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println(num);

	//Declare a pointer to an integer
	var ptr *int
	ptr = &num;

	//Pointer use to take the variable memory address
	fmt.Println(ptr);

	//That use to dereference the pointer.then take the variable value
	fmt.Println(*ptr);

	// then we can change the variable value using the pointer
	
	num = 100;
	var pointer = &num;
	*pointer = 200;

	fmt.Println(num);
	fmt.Println(&pointer);


}