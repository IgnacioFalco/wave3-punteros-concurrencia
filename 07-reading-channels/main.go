package main

import (
	"fmt"
	"time"
)

func procesarConCanal(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")

	c <- i
}

func main() {

	canal := make(chan int)

	go procesarConCanal(1, canal)

	variable := <-canal // recibimos y lo asignamos a una variable

	fmt.Println("Termino el programa ", variable)
	// fmt.Println("Que pasa si vuelvo a leer de un canal ya escrito? ", variable)

}
