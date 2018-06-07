package generator

import "fmt"

func ExampleNewIntGenerator() {
    gen := NewVertexIDGenerator()
    fmt.Println(gen(), gen(), gen())
    fmt.Println(gen(), gen(), gen())

    // Output:
    // 1 2 3
    // 4 5 6
}

func ExampleNewIntMultipleGenerator() {
    gen1 := NewVertexIDGenerator()
    gen2 := NewEdgeIDGenerator()

    fmt.Println(gen1(), gen1(), gen1())
    fmt.Println(gen2(), gen2(), gen2())
    fmt.Println(gen1(), gen1(), gen1())

    // Output:
    // 1 2 3
    // 1 2 3
    // 4 5 6
}
