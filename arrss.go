package main

import (
	"fmt"
	"time"
)

// 练习7.1
func test16() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原数组a：", a)

	var b = a
	b[3] = 8
	fmt.Println("现数组a：", a)
	fmt.Println("现数组b：", b)
}

// 练习7.2
func test17() {
	a := [15]int{}
	for i := 0; i < 15; i++ {
		a[i] = i
	}
	fmt.Println(a)
}

// 练习7.3
func test18() {
	tb := time.Now()
	last1, last2 := 0, 1
	fmt.Print("1 ")
	for i := 1; i < 50; i++ {
		var addd = last1 + last2
		last1 = last2
		last2 = addd
		fmt.Print(addd, " ")
	}

	te := time.Now()
	fmt.Println()
	fmt.Println("两个变量用时: ", te.Sub(tb))

	tb = time.Now()
	arr := [50]int{1, 1}
	fmt.Print("1 1 ")
	for i := 2; i < 50; i++ {
		arr[i] = arr[i-1] + arr[i-2]
		fmt.Print(arr[i], " ")
	}

	te = time.Now()
	fmt.Println()
	fmt.Println("数组用时: ", te.Sub(tb))
}

func test19help(a *[3]int) func() {
	return func() {
		fmt.Println(a)
	}
}
func test19aaa() func() {
	aaa := [3]int{1, 2, 3}
	return test19help(&aaa)
}
func test19() {
	f := test19aaa()
	f()
}
func test20Append(sl []byte, aaa []byte) (res []byte) {
	slLen := len(sl)
	if cap(sl) >= slLen+len(aaa) {
		res = sl[:slLen+len(aaa)]
	} else {
		res = make([]byte, slLen+len(aaa))
		for i := 0; i < len(sl); i++ {
			res[i] = sl[i]
		}
	}

	for i := 0; i < len(aaa); i++ {
		res[slLen+i] = aaa[i]
	}
	return
}

// 练习 7.5
func test20() {
	sl := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	fmt.Println(test20Append(sl, []byte{'a', 'b', 'c'}))

	fmt.Println(test20Append(test20Append(sl, []byte{'a', 'b', 'c'}), []byte{'a', 'b', 'c'}))
}
