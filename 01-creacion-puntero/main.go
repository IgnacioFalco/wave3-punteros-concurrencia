package main

import "fmt"

func main() {
	var p1 *int
	fmt.Printf("%p\n", p1)

	// Forma 2
	/* var p2 = new(int)
	fmt.Printf("%p\n", p2) */

	//Forma 3 y operador de direccion
	/* v := 12
	p3 := &v
	fmt.Printf("%p\n", p3) */

}
