package main

var x = 0

func f() bool {
	x++
	return x < 5
}

func main(){
	for f(); f(); f() {
		println(x)
	}
}