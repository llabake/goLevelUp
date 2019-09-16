package main

import "fmt"

// defining a struct type
type Developer struct {
	FirstName string
	LastName  string
	DLevel       int
}

func main() {
	// Declaring a variable of a `struct` type
	var engineer Developer // // All the struct fields are initialized with their zero value
	fmt.Println(engineer)

	// Declaring and initializing a struct using a struct literal
	D2 := Developer{"Idrees", "Ibrahim", 2}
	fmt.Println("D2 Engineer: ", D2)

	// Naming fields while initializing a struct
	D3 := Developer{
		FirstName: "Gilbert",
		LastName:  "Nzuwera",
		DLevel:       3,
	}
	fmt.Println("D3 Engineer: ", D3)

	// Uninitialized fields are set to their corresponding zero-value
	D4 := Developer{FirstName: "Edward"}
	fmt.Println("D4 Engineer: ", D4)

	// Anonymous Struct
	myPhone := struct {
		model, manufacturer, color string
		serialNumber float64
	}{
		model: "No model name now",
		manufacturer: "Infinix",
		serialNumber: 283888284884.04004,
	}
	fmt.Println("D1 Engineer Phone: ", myPhone)

	// Accessing struct fields using the dot operator
	fmt.Println("New Engineer: ", D3.FirstName)

	// Assigning a new value to a struct field
	D2.DLevel = 3
	fmt.Println("Newly Promoted: ", D2)
}