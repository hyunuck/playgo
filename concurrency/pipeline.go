package main

func PlusOne(in <- chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for num := range in {
            out <- num + 1
        }
    }()
    return out
}

