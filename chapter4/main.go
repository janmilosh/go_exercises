package main

import "fmt"


func main() {
    fmt.Print("Enter a temperature in Farenheit: ")
    var tempF float64
    fmt.Scanf("%f", &tempF)
    tempC := (tempF - 32) * 5/9
    fmt.Println(tempC)

    fmt.Print("Enter a distance in feet: ")
    var distFeet float64
    fmt.Scanf("%f", &distFeet)
    distMeter := distFeet * 0.3048
    fmt.Println(distMeter)
}

