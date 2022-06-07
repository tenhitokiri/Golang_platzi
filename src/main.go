package main

import (
	"fmt"
	"log"
	"strconv"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgs(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleValue(a int) (c, d int) {
	return a * 2, a*2 + 1
}

func main() {
	// This is a constant
	/*
		const pi float64 = 3.14
		const pi2 = 3.1415

		// This is an integer variable
		var x int = 10
		y := 50
		var z int
		z = x + y

		// This is a zero value variable
		var a int
		var b float64
		var c string
		var d bool
		var e struct{}
		//var f func()

		fmt.Println("pi: ", pi)
		fmt.Println("pi2: ", pi2)
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)
		fmt.Println("z: ", z)
		fmt.Println(a, b, c, d, e)

		const base = 10
		area := base * base
		fmt.Println("area: ", area)
		z = y - x
		fmt.Println("z: ", z)

		//Println is a function that prints the message
		helloMessage := "Hello"
		   	worldMessage := "World !"
		   	fmt.Println(helloMessage, worldMessage)
		   	fmt.Println(helloMessage, worldMessage)

		//Printf is a function that prints the message with a format
		nombre := "Platzi"
		cursos := 500
		fmt.Printf("Hola %s, tiene mas de %d cursos\n", nombre, cursos)
		fmt.Printf("Hola %v, tiene mas de %v cursos\n", nombre, cursos)

		//Spintf is a function that prints the message with a format
		message := fmt.Sprintf("Hola %s, tienes ma de %d cursos", nombre, cursos)
		fmt.Println(message)
		fmt.Printf("Nombre %T, cursos: %T \n", nombre, cursos)

		//Function that returns a value
		normalFunction("Sup doc")
		tripleArgs(1, 2, "heya")
		fmt.Println(returnValue(2))
		value1, value2 := doubleValue(2)
		fmt.Printf("value1: %d \n", value1)

		//For loop
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		//For while loop
		counter := 0
		for counter < 10 {
			fmt.Println(counter)
			counter++
		}

		//For forever loop
		counter = 0
		for {
			fmt.Println(counter)
			counter++
			if counter == 100 {
				break
			}
		}

		if value1 > value2 {
			fmt.Println("value1 is greater than value2")
			} else if value1 == value2 {
				fmt.Println("value1 is equal to value2")
				} else {
					fmt.Println("value1 is less than value2")
				}

	*/
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println(value)

	//Switch case
	/* modulo := value % 2
	switch modulo {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	default:
		fmt.Println("Error")
	}

	switch {
	case value < 10:
		fmt.Println("Less than 10")
	case value < 20:
		fmt.Println("Less than 20")
	case value < 30:
		fmt.Println("Less than 30")
	default:
		fmt.Println("Greater than 30")
	}
	*/

	//Defer
	defer fmt.Println("This is a defer")
	fmt.Println("This is a normal message")

	//Continue and Break
	/* for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println(i)
	} */

	//Arrays
	var array [5]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Printf("Array length: %d\n", len(array))
	fmt.Printf("Array capacity: %d\n", cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Slice capacity: %d\n", cap(slice))
	fmt.Printf("Slice length: %d\n", len(slice))

	//Slice methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append to slice
	slice = append(slice, 9)
	fmt.Println(slice)

	newSlide := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	slice = append(slice, newSlide...)
	fmt.Println(slice)

}
