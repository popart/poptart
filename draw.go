package main

import (
    "fmt"
    "log"
    "image"
    "image/color"
    "image/png"
    "os"
    "poptart/math"
)

func main() {
    fmt.Printf("Main Prog:\n")
    math.Test()

    img := image.NewGray(image.Rect(0, 0, 10, 20))
    img.Set(2, 2, color.Gray{0xFF})

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
