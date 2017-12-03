package render

import (
    "image"
    "image/color"
    m "math"
    "poptart/math"
)

type Gray struct {
    image.Gray
}

func NewGray(r image.Rectangle) *Gray {
    return &Gray{*image.NewGray(r)}
}

func (img *Gray) RenderPoint(p math.Point) {
    pixX := int(p.X)
    pixY := int(p.Y)
    curColor := img.GrayAt(pixX, pixY)
    img.Set(pixX, pixY, color.Gray{curColor.Y + 0x3F})
}

func (img *Gray) RenderTriangle(t math.Triangle) {
    minX := t[0].X
    maxX := t[0].X
    minY := t[0].Y
    maxY := t[0].Y

    for _, p := range t {
        minX = m.Min(minX, p.X)
        maxX = m.Max(maxX, p.X)
        minY = m.Min(minY, p.Y)
        maxY = m.Max(maxY, p.Y)
    }

    var testP math.Point
    for i := minX; i <= maxX; i+=.5 {
        for j := minY; j <= maxY; j+=.5 {
            testP = math.Point{i,j}
            if math.InsideTriangle(t, testP) {
                img.RenderPoint(testP)
            }
        }
    }
}
