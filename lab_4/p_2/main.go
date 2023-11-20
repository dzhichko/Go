package main

//go:generate go build -o gen.exe D:\GO\projects\Go_course\lab_4\p_2
//go:generate go build -o gen_pro.exe -tags pro D:\GO\projects\Go_course\lab_4\p_2

import "fmt"

var n int = 7

func main() {
	var res int = 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	fmt.Printf("Result of calculating <%d!> : %d", n, res)
}
