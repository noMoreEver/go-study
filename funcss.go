package main

import (
	"fmt"
	"math"
	"math/big"
)

// 练习6.1
func test8(a, b int) (add, multi, remove int) {
	add = a + b
	multi = a * b
	remove = a - b
	fmt.Println(add, multi, remove)
	return
	// return a + b, a * b, a - b
}

// 练习6.2
func test9(x float64) (res float64) {
	res = math.Sqrt(x)
	fmt.Println(res)
	return
}

// 练习6.3
func test10(s ...string) {
	for i := range s {
		fmt.Println(s[i])
	}
}

// test
func test11help() {
	fmt.Println("enter help")
	defer fmt.Println("defer run")
	fmt.Println("exit help")
}
func test11() {
	test11help()
	fmt.Println("back test11")
}

// 练习6.4
func test12help(a int) (t, v int) {
	t = a
	if a == 0 || a == 1 {
		v = 1
	} else {
		_, v1 := test12help(a - 1)
		_, v2 := test12help(a - 2)
		v = v1 + v2
	}
	return
}

func test12() {
	fmt.Println(test12help(10))
}

// 练习6.5
func test13(a int) {
	if a <= 0 {
		fmt.Println()
		return
	}
	fmt.Print(a, " ")
	test13(a - 1)
}

// 练习6.6
func test14(a int) *big.Int {
	bigI := big.NewInt(1)
	for i := 1; i <= a; i++ {
		bigI.Mul(bigI, big.NewInt(int64(i)))
	}
	fmt.Printf("%d的阶乘是%v\n", a, bigI)
	return bigI
}

// 练习6.7
func test15(s string) (res string) {
	for _, c := range s {
		if c > 255 {
			res += "?"
		} else {
			res += string(c)
		}
	}
	fmt.Printf("%s -> %s\n", s, res)
	return
}
