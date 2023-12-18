package main

import (
	"fmt"
)

//array are passed by value
//slice are passed by referenc
//map passed by value but can be passed as reference by pointers

func do(b []int) int {
	b[0] = 0
	return b[1]
}

func do_array(b [3]int) int {
	b[0] = 0
	return b[1]
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func prev_main() {
	a := []int{1, 2, 3}
	v := do(a)
	fmt.Println(a, v)

	a_a := [3]int{1, 2, 3}
	v_v := do_array(a_a)
	fmt.Println(a_a, v_v)

}

func main() {
	f := fib()
	for i := f(); i < 100; i = f() {
		fmt.Println(i)
	}
}

//if inside var referenced outside, then it lives on unline C/c++
