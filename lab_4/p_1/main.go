package main

func main() {
	var n int = 7

	fact(n)
}

func fact(n int) {
	var res int = 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	println(res)
}
