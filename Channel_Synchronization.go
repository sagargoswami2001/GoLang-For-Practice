package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("Sagar ")
    time.Sleep(time.Second)
    fmt.Println("Goswami")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}
