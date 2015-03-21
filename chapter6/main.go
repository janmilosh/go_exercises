package main

import "fmt"

func main() {
    x := []float64{98, 93, 77, 82, 83,}

    total := 0.0
    for _, value := range x {
        total += value
    }
    fmt.Println(total / float64(len(x)))
    fmt.Println(x[3])
    slice1 := []int{1, 2, 3,}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)

    elements := make(map[string]string)

    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"
    

    if name, ok := elements["Un"]; ok {    
        fmt.Println(name, ok)
    }

    fmt.Println(len(make([]int, 3, 9)))

    y := [6]string{"a","b","c","d","e","f"}
    fmt.Println(y[2:5])

    numbers := []int64 {
        -456, 48,96,86,68,
        57,82,63,70, 1000,
        37,34,83,27, -333,
        19,97, 9,17, 2001,
        -400, -9999999999,
    }

    smallest := numbers[0]
    for _, number := range numbers {
        if number < smallest {
            smallest = number
        }
    }
    fmt.Println("The smallest number is: ", smallest)
}