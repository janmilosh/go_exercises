package main

import "fmt"

func half(value int) (int, bool) {
    half := value/2
    return half, even(value)
}

func even(value int) bool {
    if value % 2 == 0 {
        return true
    } 
    return false
}



func main() {
    fmt.Println(half(567))
}