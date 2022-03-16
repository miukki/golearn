package main

import (
        "math/rand"
        "bufio"
        "fmt"
        "os"
        "strings"
)

const (
        supplyTotal = 20000
        numTotal = 1000
)
func main() {
        var intArr [supplyTotal/numTotal]uint16

        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Just Press Enter")
        fmt.Println("---------------------")
        intArr[0]=4 //first 4 root nft not for sale.
        for {
                fmt.Print("-> ")
                text, _ := reader.ReadString('\n')
                // convert CRLF to LF
                text = strings.Replace(text, "\n", "", -1)

                //gold supply = ID 1-200, silver supply = 201-2000, bronze = 2001-20000
                randnum := uint8(rand.Intn(255))

                randval := randnum%uint8(supplyTotal/numTotal)

//              fmt.Printf("bbb %v = %v = %d\n",randnum,(randnum % 9),randval)
                if randval == 0 {
                        if ((randnum % 9) == 1) {
                                if intArr[randval] == 96 { //so gold ticket cap at 96 + 4 root nft, remainder 100 for sale in 2nd batch sale
                                        randval = (randnum%uint8((supplyTotal/numTotal)-1))+1
                                }
                        }else{
                                randval = (randnum%uint8((supplyTotal/numTotal)-1))+1
                        }
                }
                intArr[randval] = intArr[randval]+1
                getval := intArr[randval]

                if randval == 0 {
                        fmt.Printf("CONGRATS! You got GOLD TICKET!!! ")
                }
                fmt.Printf("The #ID You Bought: %d\n",uint16(getval)+(uint16(randval)*numTotal))

        }

}