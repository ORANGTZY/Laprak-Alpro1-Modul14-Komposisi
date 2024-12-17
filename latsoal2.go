package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n) 
    if n <= 1 {
        fmt.Println("Bukan Prima") 
    } else {
        Prima := true
        for i := 2; i < n; i++ {
            if n%i == 0 {
                Prima = false
                break
            }
        }
        if Prima {
            fmt.Println("Prima")
        } else {
            fmt.Println("Bukan Prima")
        }
    }
}