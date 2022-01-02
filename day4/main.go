package main

import (
	"fmt"
	"os"
)

func main() {
	var min, max int
	fmt.Fscanf(os.Stdin, "%d-%d", &min, &max)
	var passwords []string
	for i := min; i <= max; i++ {
		passwords = append(passwords, fmt.Sprint(i))
	}

	passwords = filter(passwords, func(password string) bool {
		for i := 1; i < len(password); i++ {
			if password[i-1] > password[i] {
				return false
			}
		}

		return true
	})

	passwords = filter(passwords, func(password string) bool {
		for i := 1; i < len(password); i++ {
			if password[i-1] == password[i] {
				return true
			}
		}

		return false
	})

	fmt.Println("Part 1:", len(passwords))
	passwords = filter(passwords, func(password string) bool {
		for i := 1; i < len(password); i++ {
			if password[i-1] == password[i] && (i < 2 || password[i-2] != password[i]) && (i+1 >= len(password) || password[i] != password[i+1]) {
				return true
			}
		}

		return false
	})

	fmt.Println("Part 2:", len(passwords))
}

func filter(passwords []string, filterFunc func(string) bool) []string {
	var valid []string
	for _, password := range passwords {
		if filterFunc(password) {
			valid = append(valid, password)
		}
	}

	return valid
}
