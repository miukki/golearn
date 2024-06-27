package main

import "fmt"

type CustomMapType[T comparable, V int | string] map[T]V

func main() {

	//CustomMapType
	myMap := CustomMapType[int, string]{
		5: "5",
		6: "3",
	}

	m := make(CustomMapType[int, int])
	m[5] = 5

	for key, value := range myMap {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	for key, value := range m {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
	//CustomMapType

}
