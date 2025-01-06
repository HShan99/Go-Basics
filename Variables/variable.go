package main

import "fmt"

func main(){
	
	var firstName int = 20;
	var secondName int = 30;

	fmt.Println(firstName + secondName);
	fmt.Println(firstName - secondName);
	fmt.Println(firstName * secondName);
	fmt.Println(firstName / secondName);


	var userName string = "Shan Harhsa";
	fmt.Println(userName);

	var userScore  float32 = 23.9;
	var userBalance float64 = -2.5346466;

	fmt.Println(userScore,userBalance);
	fmt.Printf("Type: %T\n" , userBalance);

	var isUserFound bool = false;
	println(isUserFound);

	var fName, lName string= "Shan", "Rathnayake";
	fmt.Println(fName,lName);

	// That is the method use in programming
	newUser := "Hasitha";
	fmt.Println(newUser); 

	// const variable
	// Area of the circle

	var radius float32 = 10;
	const pie = 3.14;
	A := pie *(radius*radius);
	fmt.Println("Area is : " ,A);
	
}