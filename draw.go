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

    img := image.NewGray(image.Rect(0, 0, 10, 20))
    render.RenderPoint(img, 5, 10)

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
