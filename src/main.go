package main

import (
	"fmt"
	"math"
)

func main() {

	// Declaro Constantes

	const perrito string = "Olix wuff: "
	const pi float64 = 3.14

	fmt.Println(pi)
	fmt.Print(perrito)
	fmt.Println("Hola Olix es un perrito siempre le digo : ")

	base := 12
	var altura int = 14
	var area int

	fmt.Print(base, altura, area)

	// Zero Values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// area de un cuadrado base * altura

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Base del cuadrado ", areaCuadrado)

	// Operadores aritmeticos

	x := 10
	y := 50

	result := x + y

	fmt.Println("Suma: ", result)

	result = y - x

	fmt.Println("Resta: ", result)

	result = y * x

	fmt.Println("multiplicación: ", result)

	result = y / x

	fmt.Println("División: ", result)

	result = y % x

	fmt.Println("Modulo: ", result)

	//Incremental

	x++
	fmt.Println("Incremental: ", x)

	x--
	fmt.Println("Decremento: ", x)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	var high = 5
	var width = 6
	var result1 = high * width
	fmt.Println("Rectangle area is:", result1)
	// Trapezoid
	var baseA = 38.0
	var baseB = 18.0
	var high2 = 7.0
	var result2 = ((baseA + baseB) / 2.0) * high2
	fmt.Println("Trapezoid area is:", result2)
	// Circle
	var radio = 7.0
	var result3 = math.Pi * math.Pow(radio, 2)
	fmt.Println("Circle area is:", result3)

	// declaramos variables

	helloMessage := "hello"
	perritoOlix := " Olix Canson"

	fmt.Println(helloMessage, perritoOlix)
	fmt.Println(helloMessage, perritoOlix)

	fmt.Printf("%s yo tengo un perrito que se llama %s\n", helloMessage, perritoOlix)
	fmt.Printf("%v yo tengo un perrito que se llama %v\n", helloMessage, perritoOlix)

	// Sprintf

	message := fmt.Sprintf("%v yo tengo un perrito que se llama %v\n", helloMessage, perritoOlix)
	fmt.Println(message)

	//tipo de dato

	fmt.Printf("tip message %T\n", message)
	fmt.Printf("tip baseA %T\n", baseA)
	fmt.Printf("tip baseB %T\n", baseB)
	fmt.Printf("tip high2 %T\n", high2)

}
