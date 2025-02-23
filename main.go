package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println("Hello, GitHub Actions!")

    num := 5
    fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}
