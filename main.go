package main

import (
	"fmt"
	"strings"
)

// 练习4.6
func test1() {
	var s = "asSASA ddd dsjkdsjs你好 dk"
	fmt.Println(s)
	fmt.Printf("有多少字节：%d\n", len(s))
	fmt.Printf("有多少字：%d\n", len([]rune(s)))
}

// 练习5.2
func test2(m int) {
	switch m {
	case 1, 2, 3:
		fmt.Println("春季")
	case 4, 5, 6:
		fmt.Println("夏季")
	case 7, 8, 9:
		fmt.Println("秋季")
	case 10, 11, 12:
		fmt.Println("冬季")
	default:
		fmt.Println("month not exist")
	}
}

// 练习5.4
func test3() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}
	i := 0
AAA:
	fmt.Println(i)
	i++
	if i < 15 {
		goto AAA
	}

}

// 练习5.5
func test4() {
	fmt.Println("================================= tow for =============================")
	for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	fmt.Println("================================= single string =============================")
	s := ""
	for i := 0; i < 25; i++ {
		s += strings.Repeat("G", i+1) + "\n"
	}
	fmt.Printf(s)
}

// 练习5.6
func test5() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d -> %b\n", i, i)
	}
}

// 练习5.7
func test6() {
	for i := 1; i < 101; i++ {
		i3, i5 := i%3, i%5
		if i3 == 0 && i5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i3 == 0 {
			fmt.Print("    Fizz ")
		} else if i5 == 0 {
			fmt.Print("    Buzz ")
		} else {
			fmt.Printf("%8d ", i)
		}
		if i%10 == 0 {
			fmt.Println()
		}
	}
}

// 练习5.8
func test7() {
	for i := 0; i < 10; i++ {
		fmt.Println(strings.Repeat("*", 20))
	}
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	fmt.Println("Hello world!")
	test1()
	test2(2)
	test3()
	test4()
	test5()
	test6()
	test7()
	test8(10, 5)
	test9(10)
	a := [5]string{"aaa", "你好", "bbb"}
	test10(a[0:2]...)
	test11()
	b()
	test12()
	test13(10)
	test14(30)
	test15("Welcome to 中国!")
}
