package main

import (
	"fmt"
	"os"
)

func main() {
	var data string
	fmt.Fscan(os.Stdin, &data)
	width, height := 25, 6
	layers := make([]layer, len(data)/(width*height))
	for i := range layers {
		layers[i] = newLayer(width, height)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				layers[i].setValue(x, y, data[(i*height+y)*width+x]-'0')
			}
		}
	}

	minLayer, minZeroes := layers[0], layers[0].zeroes()
	for i := 1; i < len(layers); i++ {
		if zeroes := layers[i].zeroes(); zeroes < minZeroes {
			minLayer = layers[i]
			minZeroes = zeroes
		}
	}

	fmt.Println("Part 1:", minLayer.ones()*minLayer.twos())
	var image = newLayer(width, height)
	for y := 0; y < image.height; y++ {
		for x := 0; x < image.width; x++ {
			image.setValue(x, y, getValue(layers, x, y))
		}
	}

	fmt.Printf("Part 2:\n%s\n", image)
}

func getValue(layers []layer, x, y int) byte {
	for i := range layers {
		if value := layers[i].getValue(x, y); value != 2 {
			return value
		}
	}

	panic("all layers are transparent")
}
