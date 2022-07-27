/*
Create an abstraction to make conversions among different temperature systems such as Celsius, Fahrenheit and Kelvin.
Create a calculator that is loosely coupled to that abstraction
i.e. program to an interface, not an implementation principles is observed.
Make sure that the program can interact with the user so that user can achieve some conversions using the program.
*/

// Celsius = (Fahrenheit - 32) * 5 / 9
// Kelvin = Celsius + 273.15
// Fahrenheit = Celsius * 9 / 5 + 32
package main

import "fmt"

type Values struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

type Conversion interface {
	CelsiustoFahrenheit()
	CelsiustoKelvin()
	KelvintoFahrenheit()
	KelvintoCelsius()
	FahrenheittoCelsius()
	FahrenheittoKelvin()
}

//Conversion Functions
func (x Values) KelvintoFahrenheit() float64 {
	return (x.Kelvin*9/5 - 459.67)
}

func (x Values) KelvintoCelsius() float64 {
	return (x.Kelvin - 273.15)
}

func (x Values) FahrenheittoCelsius() float64 {
	return ((x.Fahrenheit - 32) * 5 / 9)
}

func (x Values) FahrenheittoKelvin() float64 {
	return ((x.Fahrenheit + 459.67) * 5 / 9)
}

func (x Values) CelsiustoFahrenheit() float64 {
	return (x.Celsius* 9/5 + 32)
}

func (x Values) CelsiustoKelvin() float64 {
	return (x.Celsius + 273.15)
}

func main() {

	//Taking values from the user
	var cels float64
	fmt.Printf("Celsius Degree :")
	fmt.Scanln(&cels)
	var fahr float64
	fmt.Printf("Fahrenheit Degree :")
	fmt.Scanln(&fahr)
	var kelv float64
    fmt.Printf("Kelvin Degree :")
	fmt.Scanln(&kelv)

	x := Values{
		Fahrenheit: fahr,
		Celsius:    cels,
		Kelvin:     kelv,
	}

	// Testing the functions
	
	fmt.Println("C to F :",x.CelsiustoFahrenheit())
	fmt.Println("C to K :",x.CelsiustoKelvin())
	fmt.Println("K to C :",x.KelvintoCelsius())
	fmt.Println("K to F :",x.KelvintoFahrenheit())
	fmt.Println("F to C :",x.FahrenheittoCelsius())
	fmt.Println("F to K :",x.FahrenheittoKelvin())
	

}
