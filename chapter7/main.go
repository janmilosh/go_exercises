package main

import "fmt"


func average(args ...float64) float64 {
    total := 0.0
    for _, v := range args {
        total += v
    }
    return total / float64(len(args)) 
}

func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }

    return x * factorial(x-1)
}

func first() {
    fmt.Println("1st")
}

func second() {
    fmt.Println("2nd")
}

func main() {
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4

    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs...))

    value := uint(10)
    s := fmt.Sprintf("The factorial of %d is %d.", value, factorial(value))
    fmt.Println(s)
    defer first()
    second()

    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC Situation Test")
}