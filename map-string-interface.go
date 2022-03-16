
package main

import (
	"fmt"
)

func main() {
	foods := map[string]interface{}{
    "bacon": "delicious",
    "eggs": struct {
      source string
      price  float64
    }{"chicken", 1.75},
    "steak": true,
    "meat": []string{"chicken", "crab", "pork"},
    "veg": []interface{}{"tomato", 1, true},
  }
  
  type Person struct {
      Name    string
      Age     int
      Hobbies []string
  }

  d := map[string]interface{}{
    "name":"John",
    "age":29,
    "hobbies": []string{"martial arts", "breakfast foods", "piano"},
  }

	fmt.Println(d)
	fmt.Println(foods)
}


