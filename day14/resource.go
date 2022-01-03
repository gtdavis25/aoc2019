package main

import "fmt"

type resource struct {
	quantity int
	chemical string
}

func (r resource) String() string {
	return fmt.Sprintf("%d %s", r.quantity, r.chemical)
}
