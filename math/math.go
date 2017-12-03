package math

import "fmt"

type Point struct {
    x, y int
}
type Vec struct {
    x, y int
}

type Triangle = [3]Point

func DotProduct(a, b Vec) int {
    return (a.x * b.x) + (a.y * b.y)
}

func Normal(a Vec) Vec {
    return Vec{-a.y, a.x}
}

func Inside(a, b, c Point) bool {
    aToB := Vec{b.x - a.x, b.y - a.y}
    aToC := Vec{c.x - a.x, c.y - a.y}

    return DotProduct(Normal(aToB), aToC) >= 0
}

func InsideTriangle(t Triangle, c Point) bool {
    return Inside(t[0], t[1], c) &&
           Inside(t[1], t[2], c) &&
           Inside(t[2], t[0], c)
}

func Test() {
    fmt.Printf("math  t\n")
}
