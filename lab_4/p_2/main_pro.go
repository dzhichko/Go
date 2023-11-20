//go:build pro
// +build pro

package main

import "fmt"

func init() {
	fmt.Print("<Pro_version == true> Enter value to calculate: ")
	_, err := fmt.Scan(&n)
	if err != nil || n < 0 {
		panic("Invalid input")
	}
}
