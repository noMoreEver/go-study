package main

import (
	"container/list"
	"fmt"
	"strconv"
)

type TreeNode struct {
	v      int
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
}

type BinTree struct {
	root *TreeNode
}

func main() {
	t := &BinTree{}
	t.Insert(8)
	fmt.Println(printTree(t.root))

	t.Insert(4)
	fmt.Println(printTree(t.root))
	t.Insert(2)
	fmt.Println(printTree(t.root))
	t.Insert(6)
	fmt.Println(printTree(t.root))
	t.Insert(1)
	fmt.Println(printTree(t.root))
	t.Insert(3)
	fmt.Println(printTree(t.root))
	t.Insert(5)
	fmt.Println(printTree(t.root))
	t.Insert(7)
	fmt.Println(printTree(t.root))
	t.Insert(12)
	fmt.Println(printTree(t.root))
	t.Insert(10)
	fmt.Println(printTree(t.root))
	t.Insert(14)
	fmt.Println(printTree(t.root))
	t.Insert(9)
	fmt.Println(printTree(t.root))
	t.Insert(11)
	fmt.Println(printTree(t.root))
	t.Insert(13)
	fmt.Println(printTree(t.root))
	t.Insert(15)
	fmt.Println(printTree(t.root))

	t.Remove(1)
	fmt.Println(printTree(t.root))
	t.Remove(2)
	fmt.Println(printTree(t.root))
	t.Remove(4)
	fmt.Println(printTree(t.root))
}

// 层序遍历
func (root *TreeNode) level1by1() []int {
	res := make([]int, 10)
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		n := l.Remove(l.Front()).(*TreeNode)
		res = append(res, n.v)
		if n.left != nil {
			l.PushBack(n.left)
		}
		if n.right != nil {
			l.PushBack(n.right)
		}
	}
	return res
}

// 深度遍历 -- 前序
func (root *TreeNode) depth1by1() []int {
	res := make([]int, 10)
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		n := l.Remove(l.Back()).(*TreeNode)
		res = append(res, n.v)
		if n.right != nil {
			l.PushBack(n.right)
		}
		if n.left != nil {
			l.PushBack(n.left)
		}
	}

	return res
}

var res1 = make([]int, 10) // 全局变量，在并行时不安全

// 深度遍历 -- 中序
func (node *TreeNode) depth1by1Center() {
	if node == nil {
		return
	}
	node.left.depth1by1Center()
	res1 = append(res1, node.v)
	node.right.depth1by1Center()
	// res1 = append(res1, node.v)

}

// TODO: 以栈的形式完成中序和后序深度遍历

// 二叉树查找节点
func (t *BinTree) search(v int) *TreeNode {
	if t.root == nil {
		return nil
	}
	n := t.root
	for {
		if n.v > v {
			if n.left == nil {
				return n
			}
			n = n.left
		} else if n.v < v {
			if n.right == nil {
				return n
			}
			n = n.right
		} else if n.v == v {
			return n
		}
	}
}

func (t *BinTree) Insert(v int) {
	n := t.search(v)
	if n == nil { // 无根节点
		t.root = &TreeNode{v: v}
	} else if n.v == v {
		return // 重复元素无需插入
	} else if n.v > v {
		n.left = &TreeNode{v: v}
		n.left.parent = n
	} else if n.v < v {
		n.right = &TreeNode{v: v}
		n.right.parent = n
	}
}

func (t *BinTree) Remove(v int) {
	n := t.search(v)
	fmt.Println(n)
	// 元素不存在，直接退出
	if n == nil || n.v != v {
		return
	}
	// 节点为根节点
	if n.parent == nil {
		t.root = nil
	}
	lor := true
	if n.parent.right != nil && n.v == n.parent.right.v {
		lor = false
	}
	// 子节点数为0
	if n.left == nil && n.right == nil {
		if lor {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	} else if n.left != nil && n.right != nil { // 子节点数为2
		tmp := n.right
		for tmp.left != nil {
			tmp = tmp.left
		}
		t.Remove(tmp.v)
		n.v = tmp.v
	} else if n.left != nil {
		if lor {
			n.parent.left = n.left
		} else {
			n.parent.right = n.left
		}
	} else if n.right != nil {
		if lor {
			n.parent.left = n.right
		} else {
			n.parent.right = n.right
		}
	}
}

// https://pkg.go.dev/github.com/imgoogege/LeetCode-in-Go/Algorithms/0655.print-binary-tree#section-readme
func printTree(root *TreeNode) [][]string {
	level := getLevel(root)
	size := 1<<uint(level) - 1
	res := make([][]string, level)
	for i := range res {
		res[i] = make([]string, size)
	}

	// loc 是 root　节点的值在 res 中的索引号
	loc := 1<<uint(level-1) - 1
	getRes(root, 0, loc, res)

	return res
}

func getRes(root *TreeNode, i, j int, res [][]string) {
	if root == nil {
		return
	}

	res[i][j] = strconv.Itoa(root.v)

	level := len(res)
	if level-i-2 < 0 {
		return
	}

	diff := 1 << uint(level-i-2)

	getRes(root.left, i+1, j-diff, res)
	getRes(root.right, i+1, j+diff, res)
}

func getLevel(root *TreeNode) (res int) {
	res = 1
	queue := make([]*TreeNode, 1, 1024)
	queue[0] = root

	for {
		hasChild := false
		for i := 0; i < len(queue); i++ {
			if queue[i].left != nil || queue[i].right != nil {
				hasChild = true
				break
			}
		}

		if !hasChild {
			break
		}
		res++

		fQueue := queue[:len(queue)]
		queue = queue[len(queue):]

		for i := 0; i < len(fQueue); i++ {
			if fQueue[i].left != nil {
				queue = append(queue, fQueue[i].left)
			}
			if fQueue[i].right != nil {
				queue = append(queue, fQueue[i].right)
			}
		}
	}

	return res
}
