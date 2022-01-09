package main

import "fmt"

type deck struct {
	size, increment, offset int
}

func newDeck(size int) *deck {
	return &deck{
		size:      size,
		increment: 1,
	}
}

func (d *deck) card(i int) int {
	return mod(d.increment*i+d.offset, d.size)
}

func (d *deck) cut(n int) {
	d.offset = mod(d.offset+n*d.increment, d.size)
}

func (d *deck) dealWithIncrement(n int) {
	d.increment = modMultiply(d.increment, modInverse(n, d.size), d.size)
}

func (d *deck) dealIntoNewStack() {
	d.dealWithIncrement(d.size - 1)
	d.cut(1)
}

func (d *deck) shuffle(steps []string) {
	for _, line := range steps {
		var n int
		if line == "deal into new stack" {
			d.dealIntoNewStack()
		} else if _, err := fmt.Sscanf(line, "cut %d", &n); err == nil {
			d.cut(n)
		} else if _, err := fmt.Sscanf(line, "deal with increment %d", &n); err == nil {
			d.dealWithIncrement(n)
		} else {
			panic("invalid step: " + line)
		}
	}
}

func (d *deck) composeWith(other *deck) *deck {
	increment := modMultiply(d.increment, other.increment, d.size)
	offset := mod(modMultiply(d.increment, other.offset, d.size)+d.offset, d.size)
	return &deck{
		size:      d.size,
		increment: increment,
		offset:    offset,
	}
}

func (d *deck) iterate(times int) *deck {
	acc := &deck{increment: 1, offset: 0, size: d.size}
	t := d
	for ; times > 0; times >>= 1 {
		if times&1 == 1 {
			acc = acc.composeWith(t)
		}

		t = t.composeWith(t)
	}

	return acc
}

func modMultiply(m, n, p int) int {
	x, y := 0, m
	for t := n; t != 0; t >>= 1 {
		if t&1 == 1 {
			x = mod(x+y, p)
		}

		y = mod(y+y, p)
	}

	return x
}

func modInverse(n, p int) int {
	d, x, _ := extendedEuclid(n, p)
	if d != 1 {
		panic(fmt.Sprintf("%d is not invertible mod %d", n, p))
	}

	return mod(x, p)
}

func extendedEuclid(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}

	d, x, y := extendedEuclid(b, a%b)
	return d, y, x - a/b*y
}

func mod(n, p int) int {
	r := n % p
	if r < 0 {
		return p + r
	}

	return r
}
