package main

func main() {
	var a int

	var (
		b int
		c bool
	)

	a = 1
	b = 2
	c = true
	d := 4
	a = a + b + d
	c = c && false

	//explicit declaration of constants
	const (
		e        = 0
		f string = "0"
	)
}
