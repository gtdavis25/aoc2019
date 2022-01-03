package main

import (
	"fmt"
	"strings"
)

type reaction struct {
	inputs []resource
	output resource
}

func (r reaction) String() string {
	w := new(strings.Builder)
	for i := range r.inputs {
		if i > 1 {
			fmt.Fprint(w, ", ")
		}

		fmt.Fprint(w, r.inputs[i])
	}

	fmt.Fprintf(w, " => %s", r.output)
	return w.String()
}
