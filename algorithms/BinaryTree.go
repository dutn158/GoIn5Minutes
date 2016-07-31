package main

import "fmt"

type NODE struct {
	key int
	left *NODE
	right *NODE
}

func createNode(x int) *NODE {
	n := NODE{
		key: x,
	}

	return &n
}

func findInsert(root *NODE, x int) *NODE {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root
	}
	if root.right == nil {
		return root
	}
	if root.key > x {
		return findInsert(root.left, x)
	} else if root.key < x {
		return findInsert(root.right, x)
	}
	return nil
}

func insertNode(root *NODE, x int) {
	create := createNode(x)
	if root == nil {
		root = create
		return
	}
	insert := findInsert(root, x)

	if insert.key > x {
		root.left = create
	} else if insert.key < x {
		root.right = create
	}
}

func createTree(root *NODE, a []int, n int) {
	for i := 0; i < n; i++ {
		insertNode(root, a[i])
	}
}

func LNR(root *NODE) {
	if root != nil {
		LNR(root.left)
		fmt.Printf("%v\t", root.key)
		LNR(root.right)
	}
}

func main() {
	var root *NODE
	a := []int{43, 21, 9, 31, 35, 10, 19, 15, 65, 56, 62}
	createTree(root, a, len(a))
	LNR(root)
}
