package main

type moveResult int

const (
	hitWall moveResult = iota
	moved
	found
)
