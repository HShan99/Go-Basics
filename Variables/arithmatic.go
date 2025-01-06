package main

import "fmt"

func main() {

	a := 10
	b := 3

	//Arithmetics operation
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Println("Sum : ",sum);
	fmt.Println("Difference : ",difference);
	fmt.Println("Product : ", product);
	fmt.Println("Quotient : ",quotient);
	fmt.Println("Remainder : ", remainder);

	// Also use that method 
	a += 2
	fmt.Println(a);

	a -= 5
	fmt.Println(a);

	a *= 3
	fmt.Println(a);

	a /= 2
	fmt.Println(a);

	// if you want to operations with 1 we can ue increment,decrement operators
	a++;
	fmt.Println(a);	

	a--;
	fmt.Println(a);

	// Bitwise operator
	fmt.Println(5&3);
	fmt.Println(5|3);

}