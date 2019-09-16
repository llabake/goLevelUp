package main

import "fmt"

func main () {
	// a variable being declared at the point of initialization
	name := "Justina"
	stack := "Javascript/React"
	
	// a pointer that points to the memory location of the declared varaible
	pname := &name

	fmt.Println("My name is ", name, "and I am decalared in the memory location", pname,
	" \n I can also get my value back using ", *pname, "this is called dereferencing")


	fmt.Println("Hi ", name, "your current stack is ", stack)

	changeStack(&stack)

	fmt.Println(name, "you are now on the ", stack)
}

func changeStack(stack *string) {
	*stack = "Python/Django"

	fmt.Println("Are you sure you want to change stack to ", *stack, "?")
	return *stack
}