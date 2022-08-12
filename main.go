package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    var count int
    if len(os.Args) < 2 { return }
    count = len(strings.Fields(os.Args[1]))
    fmt.Println(count)
}
