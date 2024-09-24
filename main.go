package main

import (
    "fmt"
    "time"
    "github.com/go-vgo/robotgo"
    "math/rand"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    for {
        x, y := robotgo.GetMousePos()
        newX := x + rand.Intn(10) - 5
        newY := y + rand.Intn(10) - 5

        robotgo.MoveMouse(newX, newY)
        fmt.Println("Mouse moved to", newX, newY)
        time.Sleep(5 * time.Second)
    }
}