package render

import (
    "image"
    "image/color"
)

type Gray struct {
    image.Gray
}

func NewGray(r image.Rectangle) *Gray {
    return &Gray{*image.NewGray(r)}
}

func (img *Gray) RenderPoint(x, y int) {
    img.Set(x, y, color.Gray{0xFF})
}
