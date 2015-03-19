package main

import "fmt"

func main() {
    var name string

    for i := 1; i <= 10; i++ {
        switch i {
        case 1: name = "One"
        case 2: name = "Two"
        case 3: name = "Three"
        case 4: name = "Four"
        case 5: name = "Five"
        case 6: name = "Six"
        case 7: name = "Seven"
        case 8: name = "Eight"
        case 9: name = "Nine"
        case 10: name = "Ten"
        default: name = "Who knows"
        }

        if i % 2 == 0 {
            fmt.Println(name, "is even.")
        } else {
            fmt.Println(name, "is odd.")
        }
    }

    i := 10
    if i > 10 {
        fmt.Println("Big")
    } else {
        fmt.Println("Small")
    }

    for i := 1; i <= 100; i++ {
        if i % 3 == 0 {
            fmt.Println(i)
        }
    }

    for i := 1; i <= 100; i++ {
        if i % 15 == 0 {
            fmt.Println(i, "FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println(i, "Fizz")
        } else if i % 5 == 0 {
            fmt.Println(i, "Buzz")
        } else {
            fmt.Println(i)
        }
    }

    for i := 1; i <= 100; i++ {
        switch {
        case (i % 15 == 0): fmt.Println(i, "FizzBuzz")
        case (i % 3 == 0): fmt.Println(i, "Fizz")
        case (i % 5 == 0): fmt.Println(i, "Buzz")
        default: fmt.Println(i)
        }        
    }
}