package render

import (
    "image"
    "image/color"
    "poptart/math"
)

type Gray struct {
    image.Gray
}

func NewGray(r image.Rectangle) *Gray {
    return &Gray{*image.NewGray(r)}
}

func (img *Gray) RenderPoint(p math.Point) {
    img.Set(p.X, p.Y, color.Gray{0xFF})
}

func (img *Gray) RenderTriangle(t math.Triangle) {
    minX := t[0].X
    maxX := t[0].X
    minY := t[0].Y
    maxY := t[0].Y

    for _, p := range t {
        minX = math.Min(minX, p.X)
        maxX = math.Max(maxX, p.X)
        minY = math.Min(minY, p.Y)
        maxY = math.Max(maxY, p.Y)
    }

    var testP math.Point
    for i := minX; i <= maxX; i++ {
        for j := minY; j <= maxY; j++ {
            testP = math.Point{i,j}
            if math.InsideTriangle(t, testP) {
                img.RenderPoint(testP)
            }
        }
    }
}
