package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var lines []string
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		lines = append(lines, scanner.Text())
	}

	d := newDeck(10007)
	d.shuffle(lines)
	fmt.Println("Part 1:", findCard(d, 2019))
	d = newDeck(119315717514047)
	d.shuffle(lines)
	d = d.iterate(101741582076661)
	fmt.Println("Part 2:", d.card(2020))
}

func findCard(d *deck, card int) int {
	for i := 0; i < d.size; i++ {
		if d.card(i) == card {
			return i
		}
	}

	return -1
}
