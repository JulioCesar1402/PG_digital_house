package weather

import "fmt"

func Weather() {
	var temperature, humidity, atmPressure uint = 14, 42, 1020
	fmt.Println("Temperature:",temperature, "-", "Humidity:", humidity, "-", "AtmPressure:",atmPressure)
}