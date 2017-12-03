package render

import (
    "image"
    "image/color"
)

func RenderPoint(img *image.Gray, x, y int) {
    img.Set(x, y, color.Gray{0xFF})
}
