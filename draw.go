package main

import (
    "fmt"
    "log"
    "image"
    "image/png"
    "os"
    "poptart/math"
    "poptart/render"
)

func main() {
    fmt.Printf("Main Prog:\n")
    math.Test()

    img := render.NewGray(image.Rect(0, 0, 100, 100))
    img.RenderPoint(math.Point{5, 10})
    img.RenderTriangle(math.Triangle{math.Point{20, 20},
        math.Point{40, 30},
        math.Point{10, 50}})

    f, err := os.Create("out.png")
    if err != nil {
        log.Fatal(err)
    }

    err = png.Encode(f, img)
    if err != nil {
        log.Fatal(err)
    }

    err = f.Close()
    if err != nil {
        log.Fatal(err)
    }
}
