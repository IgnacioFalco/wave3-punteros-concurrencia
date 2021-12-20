package main

import (
	"fmt"
	"time"
)

func procesar(i int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
}

func main() {
	for i := 0; i < 5; i++ {
		procesar(i)
	}

	// tiempo de espera para mostrar la ejecucion de todos
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Termino el programa")

}
