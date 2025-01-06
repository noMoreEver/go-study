package main

import "fmt"

type myNode struct {
	data int
	next *myNode
}
type myStack struct {
	start *myNode
	len   int
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("==================================================")
	stack := new(myStack)
	fmt.Println("stack peek: ", stack.peek())

	fmt.Println("初始stack: ", stack)
	stack.push(3)
	fmt.Println("stack 状态：", stack)
	fmt.Println("stack peek: ", stack.peek())
	stack.push(4)
	fmt.Println("stack 状态：", stack)
	stack.push(5)
	fmt.Println("stack 状态：", stack)
	fmt.Println("stack peek: ", stack.peek())
	fmt.Println("stack pop: ", stack.pop())
	fmt.Println("stack 状态：", stack)

	fmt.Println("stack pop: ", stack.pop())
	fmt.Println("stack pop: ", stack.pop())
	fmt.Println("stack pop: ", stack.pop())
	fmt.Println("stack pop: ", stack.pop())
	fmt.Println("stack 状态：", stack)

}

func (stack *myStack) push(v int) {
	newNode := &myNode{v, nil} // 不像C/C++,go指针指向局部变量，局部变量也可全局使用
	if stack.len == 0 {
		stack.start = newNode
	} else {
		newNode.next = stack.start
		stack.start = newNode
	}

	stack.len += 1
}

func (stack *myStack) pop() int {
	if stack.len == 0 {
		return 0
	} else {
		stack.len -= 1
		v := stack.start.data
		stack.start = stack.start.next
		return v
	}
}

func (stack *myStack) peek() int {
	if stack.len == 0 {
		return 0
	} else {
		return stack.start.data
	}
}

func (stack *myStack) isEmpty() bool {
	return stack.len == 0
}

func (stack *myStack) String() string {
	if stack.isEmpty() {
		return "[empty stack]"
	}
	res := "[ "
	n := stack.start
	for {
		if n == nil {
			return res + " ]"
		}
		res += fmt.Sprint(n.data) + ", "
		n = n.next
	}
}
