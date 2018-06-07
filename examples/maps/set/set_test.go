package set

import "fmt"

func Example_hasDupeRune() {
    res1 := hasDupeRune("Hello")
    res2 := hasDupeRune("world")

    fmt.Println(res1)
    fmt.Println(res2)

    // Output:
    // true
    // false
}

