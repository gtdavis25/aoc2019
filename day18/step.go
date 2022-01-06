package main

type step struct {
	position point
	previous *step
}

func (s step) length() int {
	if s.previous != nil {
		return 1 + s.previous.length()
	}

	return 0
}
