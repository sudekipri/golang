package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

type square struct {
    sideLength float64
}

type triangle struct {
    base   float64
    height float64
}

func (s square) getArea() float64 {
    return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
    return 0.5 * t.base * t.height
}

type shape interface {
    getArea() float64
}

func printArea(s shape) {
    fmt.Printf("The area of the shape is %v\n", s.getArea())
}

func promptUserForFloat(message string) (float64, error) {
    fmt.Print(message)
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        return 0, err
    }
    input = strings.TrimSpace(input)
    value, err := strconv.ParseFloat(input, 64)
    if err != nil {
        return 0, err
    }
    return value, nil
}

func main() {
    var shapeType string
    for shapeType != "square" && shapeType != "triangle" {
        fmt.Print("Enter the shape type (square/triangle): ")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        shapeType = strings.TrimSpace(input)
    }

    var s shape
    if shapeType == "square" {
        sideLength, err := promptUserForFloat("Enter the side length: ")
        if err != nil {
            fmt.Println("Error reading side length:", err)
            return
        }
        s = square{sideLength: sideLength}
    } else {
        base, err := promptUserForFloat("Enter the base length: ")
        if err != nil {
            fmt.Println("Error reading base length:", err)
            return
        }
        height, err := promptUserForFloat("Enter the height: ")
        if err != nil {
            fmt.Println("Error reading height:", err)
            return
        }
        s = triangle{base: base, height: height}
    }

    printArea(s)
}
