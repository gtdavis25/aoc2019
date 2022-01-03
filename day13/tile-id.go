package main

type tileID int

const (
	empty tileID = iota
	wall
	block
	paddle
	ball
)
