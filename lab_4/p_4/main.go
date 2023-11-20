package main

import (
	"embed"
)

//go:embed tests/file_1.txt
var fileString string

//go:embed tests/file_1.txt
var fileByte []byte

//go:embed tests/*.hash
var tests embed.FS

func main() {
	println(fileString)
	println(string(fileByte))
	content1, _ := tests.ReadFile("tests/file_2.hash")
	println(string(content1))
	content2, _ := tests.ReadFile("tests/file_3.hash")
	println(string(content2))
}
