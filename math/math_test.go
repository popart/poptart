package math

import "testing"

func TestDotProduct(t *testing.T) {
    v1 := Vec{0,1}
    v2 := Vec{1,0}
    if DotProduct(v1, v2) != 0 {
        t.Errorf("bad dot product")
    }
    if DotProduct(Vec{0,3}, Vec{1,2}) != 6 {
        t.Errorf("bad dot product")
    }
}

func TestInside(t *testing.T) {
    p1 := Point{0,0}
    p2 := Point{10,0}

    if !Inside(p1, p2, Point{1, 2}) {
        t.Errorf("bad inside")
    }
    if !Inside(p1, p2, Point{-5, 0}) {
        t.Errorf("bad inside")
    }
    if Inside(p1, p2, Point{5, -1}) {
        t.Errorf("bad inside")
    }
}

func TestInsideTriangle(t *testing.T) {
    tri := Triangle{Point{0, 0}, Point{10, 0}, Point{10, 10}}
    if !InsideTriangle(tri, Point{5, 1}) {
        t.Errorf("bad inside triangle")
    }
    if !InsideTriangle(tri, Point{5, 0}) {
        t.Errorf("bad inside triangle")
    }
    if InsideTriangle(tri, Point{5, -1}) {
        t.Errorf("bad inside triangle")
    }
    if InsideTriangle(tri, Point{5, 10}) {
        t.Errorf("bad inside triangle")
    }
}
