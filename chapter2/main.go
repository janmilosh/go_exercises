package main 

import "fmt"

// this is a comment

func main() {
    fmt.Print("Enter a distance in feet: ")
    var distanceInFeet float64
    fmt.Scanf("%f", &distanceInFeet)
    distanceInMeters := distanceInFeet * 0.3048
    fmt.Println(distanceInMeters)
}