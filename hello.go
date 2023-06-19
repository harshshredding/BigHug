package main

import "fmt"
import "rsc.io/quote/v4"

func main() {
    fmt.Println(quote.Go())
    s1 := "string1"
    s2 := "string2"
    combined := s1 + " " + s2
    fmt.Println(combined)
    var some_int = 32
    int_string := fmt.Sprintf("%d", some_int)
    fmt.Println(int_string + "\n" + combined)
}
