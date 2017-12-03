package math

import "fmt"

type Point struct {
    X, Y float64
}
type Vec struct {
    X, Y float64
}

type Triangle = [3]Point

func DotProduct(a, b Vec) float64 {
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
