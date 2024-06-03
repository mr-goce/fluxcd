package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Random number:", rand.Intn(100))
	fmt.Println("time", time.Now())
}