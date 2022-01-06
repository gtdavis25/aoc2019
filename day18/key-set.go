package main

import "strings"

type keySet int

func (k keySet) has(key tile) bool {
	return k>>(25+key-'a')&1 == 1
}

func (k keySet) add(key tile) keySet {
	return k | 1<<(25+key-'a')
}

func (k keySet) canOpen(door tile) bool {
	return k.has(door | 32)
}

func (k keySet) containsAll(other keySet) bool {
	return other & ^k == 0
}

func (k keySet) String() string {
	w := new(strings.Builder)
	for ch := byte('a'); ch <= 'z'; ch++ {
		if k.has(tile(ch)) {
			w.WriteByte(ch)
		}
	}

	return w.String()
}
