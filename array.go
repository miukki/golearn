
package main

import (
	"fmt"
)

type RandomT struct {
  one, two int
  fruit string
  truth bool
}

func main() {
	arr := []interface{}{1, 2, "apple", true}
    fmt.Println(arr)

    // interface conversion: interface {} is int
    i := arr[0].(int)
    fmt.Printf("i: %d, i type: %T\n", i, i)

    s := arr[2].(string)
    fmt.Printf("s: %s, s type: %T\n", s, s)

    _struct := struct{
      one, two int
      fruit string
      truth bool
    }{1, 2, "apple", true}
    fmt.Println(_struct)
}


