package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	pk "curso_golang_platzi/src/mypackage"
)

//Struct
type persona struct {
	nombre string
	edad   int
}

//Struct with methods
type persona2 struct {
	nombre string
	edad   int
}

func (p persona2) getNombre() string {
	return p.nombre
}

func (p *persona2) setNombre(nombre string) {
	p.nombre = nombre
}

type car struct {
	brand string
	year  int
}

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

func isPalindrome(text string) bool {
	var textReversed string
	for _, v := range text {
		textReversed = string(v) + textReversed
	}
	return strings.ToLower(textReversed) == strings.ToLower(text)
}


//Struct
type pc struct {
	ram int
	cpu float64
	brand string
	year int
	disk int
}

func (myPC pc) ping (){
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam (){
	myPC.ram = myPC.ram * 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("The PC brand is %s from year %d, has %f GHz of CPU, %d GB of RAM and %d GB of Hard Drive", myPc.brand, myPc.year, myPc.cpu, myPc.ram, myPc.disk)
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

	//Range in a Slice
	slice3 := []string{"hola", "mundo", "platzi"}
	for _, v := range slice3 {
		fmt.Println(v)
	}

	//fmt.Println(isPalindrome("ama"))
	//fmt.Println(isPalindrome("Ama"))

	//Map
	//map1 := make(map[string]int)
	//map1["Jose"] = 14
	//map1["Juan"] = 20
	//map1["Rosa"] = 32
	//fmt.Println(map1)
	//key := "Josef"
	//value, ok := map1[key]
	//if ok {
	//	fmt.Printf("%s => %d \n", key, value)
	//} else {
	//	fmt.Printf("%s Not found! \n", key)
	//}

	//delete(map1, "Jose")
	//fmt.Println(map1)

	//Range in a Map
	//for k, v := range map1 {
	//	fmt.Println(k, v)
	//}

	myCar := car{
		brand: "Mustang",
		year:  1964,
	}

	var persona1 persona
	persona1.nombre = "Juan"
	persona1.edad = 25
	fmt.Println(persona1, myCar)

	persona2 := persona2{
		nombre: "Juan",
		edad:   25,
	}

	fmt.Println(persona2)

	var myCar2 pk.CarPublic
	myCar2.Name = "Mustang"
	myCar2.Brand = "Chevy"
	myCar2.Year = 1964
	fmt.Println(myCar2)

	//no se puede acceder a los atributos de la clase privada	
	//var anotherCar pk.carPrivate
	//anotherCar.Name = "Mustang"
	//anotherCar.Brand = "Chevy"
	//anotherCar.Year = 1964
	//fmt.Println(anotherCar)

	pk.PrintMessage("Hola mundo")
	//pk.printMessage1("Hola mundo")
	a := 50
	b := &a
	
	//Pointers
	fmt.Println(a) //regular value
	fmt.Println(b) //pointer to the value of a
	fmt.Println(*b) //value of a

	*b = 51 //change the value of a with the value of brand
	fmt.Println(a) //51

	myPc := pc{
		brand: "Macbook",
		year:  2019,
		disk:  200,
		ram:   8,
		cpu:   4.5,
	}
	fmt.Println(myPc)
	
	myPc.ping()
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
}
