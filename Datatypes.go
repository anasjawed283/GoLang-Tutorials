package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Create a new scanner to read user input from standard input (console)
    scanner := bufio.NewScanner(os.Stdin)

    //enter an integer
    fmt.Print("Enter an integer: ")
    scanner.Scan() // Read user input
    inputInt, err := strconv.Atoi(scanner.Text()) // Convert input to integer
    if err != nil {
        fmt.Println("Invalid input:", err)
        return
    }

    //enter a float number
    fmt.Print("Enter a floating-point number: ")
    scanner.Scan() // Read user input
    inputFloat, err := strconv.ParseFloat(scanner.Text(), 64) // Convert input to float64
    if err != nil {
        fmt.Println("Invalid input:", err)
        return
    }

    //enter a string
    fmt.Print("Enter a string: ")
    scanner.Scan() // Read user input
    inputString := scanner.Text() // Read input as string

    fmt.Println("Integer:", inputInt)
    fmt.Println("Float:", inputFloat)
    fmt.Println("String:", inputString)
}
