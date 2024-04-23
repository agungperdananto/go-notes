package main

import (
	"fmt"
)

// early return
// guard clause

func main() {
	// declaration
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("Hello world and numbers: %v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	fmt.Println("=============================================================================")
	// short declaration
	numCars := 10
	temperature := 0.0
	var isfunny = true

	fmt.Printf("numbers of car: %v \n", numCars)
	fmt.Printf("temp: %f \n", temperature)
	fmt.Printf("is funny: %v \n", isfunny)

	fmt.Printf("temparature format is %T\n", temperature)

	// same line declaration
	mileage, company := 80987, "Tesla"

	fmt.Printf("%q has %v of mileage \n", company, mileage)

	// Default types
	// bool
	// string
	// int
	// uint32
	// byte
	// rune
	// float64
	// complex128
	fmt.Println("=============================================================================")

	const name = "Saul Goodman"
	const openRate = 30.5

	// nameLength := getLength(name)

	msg := fmt.Sprintf("Hi %v, your open rate is %.2f percent", name, openRate)

	fmt.Println(msg)

	fmt.Println("=============================================================================")

	msgLen := 10
	maxLen := 20

	fmt.Println("Trying to send message of length:", msgLen)

	if msgLen > maxLen {
		fmt.Println("message sent")
	} else {
		fmt.Println("message not sent")
	}
	fmt.Println("=============================================================================")

	if length := 10; length > 1 {
		fmt.Printf("Approved length line [%v] chars\n", length)
	}

}
