package main

import (
  "math/rand"
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main() {

  var intArr [10]uint16
  fmt.Printf("intArr %d\n",intArr)

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Just Press Enter")
  fmt.Println("---------------------")

  for {
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    randval := uint16(rand.Intn(10))
    intArr[randval] = intArr[randval]+1
    getval := intArr[randval]
	fmt.Printf("intArr %d\n",intArr)

    fmt.Printf("Your BAZI Lucky Number : %d\n",uint16(getval)+(uint16(randval)*2000))

    fmt.Printf("randval %d\n",randval)
    fmt.Printf("getval %d\n",getval)

    var i uint16=0
    for ;i<10;i++ {
        fmt.Printf("i%d = intArr[i]%d = ____ %d\n",i,intArr[i],uint16(intArr[i])+(uint16(i)*uint16(2000)))
    }
    fmt.Printf("\n")


  }

}
