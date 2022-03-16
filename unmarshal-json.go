
package main

import (
	"fmt"
)


func main() {
  var p map[string]interface{}
  err = json.Unmarshal(data, &p)
  fmt.Println(err)

  for k, v := range p {
    switch c := v.(type) {
    // case string:
    //   fmt.Printf("Item %q is a string, containing %q\n", k, c)
    // case float64:
    //   fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
    default:
      fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
    }
  }
}


