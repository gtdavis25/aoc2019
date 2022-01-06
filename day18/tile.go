package main

type tile byte

func (t tile) isKey() bool {
	return 'a' <= t && t <= 'z'
}

func (t tile) isDoor() bool {
	return 'A' <= t && t <= 'Z'
}
