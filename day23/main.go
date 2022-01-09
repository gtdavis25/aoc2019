package main

import (
	"aoc2019/intcode"
	"fmt"
	"os"
)

func main() {
	program, err := intcode.ReadProgram(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	n := newNetwork(program)
	for len(n.packets[255]) == 0 {
		n.run()
	}

	fmt.Println("Part 1:", n.packets[255][0].y)
	var natPackets []packet
	for {
		if n.idle() {
			p := n.packets[255][len(n.packets[255])-1]
			n.computers[0].Input().Write(p.x)
			n.computers[0].Input().Write(p.y)
			natPackets = append(natPackets, p)
			if len(natPackets) > 1 &&
				natPackets[len(natPackets)-2].y == natPackets[len(natPackets)-1].y {
				break
			}
		}

		n.run()
	}

	fmt.Println("Part 2:", natPackets[len(natPackets)-1].y)
}
