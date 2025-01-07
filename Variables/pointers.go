package main

import "fmt"

// increment function
func increment( value *int){
	*value ++;
}

//Set to zero function
func setToZero(value *int){
	*value = 0;
}

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

	num1 := 20;
	increment(&num1);

	fmt.Println("After increment : ", num1);

	setToZero(&num1);
	fmt.Println(num1);
	
	

}