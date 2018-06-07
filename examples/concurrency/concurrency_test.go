package concurrency

import "fmt"

func Example_simplechannel() {
    c := make(chan int)
    go func() {
        defer close(c)
        c <- 1
        c <- 2
        c <- 3
    }()

    for num := range c {
        fmt.Println(num)
    }

    // Output:
    // 1
    // 2
    // 3
}
