package main

type vertex struct {
	position point
	label    string
	edges    []edge
	isOuter  bool
}
