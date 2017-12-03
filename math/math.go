package math

import "fmt"

type Point struct {
    X, Y int
}
type Vec struct {
    X, Y int
}

type Triangle = [3]Point

func Min(x, y int) int {
    if x <= y {
        return x
    }
    return y
}

func Max(x, y int) int {
    if x >= y {
        return x
    }
    return y
}

func DotProduct(a, b Vec) int {
    return (a.X * b.X) + (a.Y * b.Y)
}

func Normal(a Vec) Vec {
    return Vec{-a.Y, a.X}
}

func Inside(a, b, c Point) bool {
    aToB := Vec{b.X - a.X, b.Y - a.Y}
    aToC := Vec{c.X - a.X, c.Y - a.Y}

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
