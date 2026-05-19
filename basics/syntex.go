package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func variable() {
	a := "This is a string."
	fmt.Println(a)

	b := 10
	fmt.Println(b)

	var c string = "This is a string."
	fmt.Println(c)
}

func maps() {
	myMap := make(map[string]int) // give a type to the key and the value

	myMap["ONE"] = 1
	myMap["TWO"] = 2

	fmt.Println(myMap)
}

func loop() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// Go has no "while"loop, still use for

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// for range loop over slices
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// infinity loop
	// for {
	// 	fmt.Println("This is an infinity loop.")
	// }
}

func conditional() {
	if 10 > 5 {
		fmt.Println("10 is greater than 5.")
	} else {
		fmt.Println("10 is not greater than 5.")
	}

	// switch statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("This is macOS.")
	case "linux":
		fmt.Println("This is Linux.")
	default:
		fmt.Printf("%s is not supported.\n", os)
	}
}

func errorHandling() {
	f, err := os.Open("file.txt")

	if err != nil { // nil is Go's zero value for pointers, interfaces, slices, maps, channels, and functions — meaning "no value" or "not initialized."
		log.Fatal(err)
	}

	n, err := f.Read(make([]byte, 1024))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes.\n", n)
}
