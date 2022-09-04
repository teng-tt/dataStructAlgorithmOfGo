package main

import "fmt"

/*
此处主要以二叉树为例
二叉树：每个节点最多只有两个儿子节点的树

二叉树的特性
高度为 h≥0 的二叉树至少有 h+1 个结点，比如最不平衡的二叉树就是退化的线性链表结构，
所有的节点都只有左儿子节点，或者所有的节点都只有右儿子节点。
高度为 h≥0 的二叉树至多有 2^h+1 个节点，比如这棵树是满二叉树。
含有 n≥1 个结点的二叉树的高度至多为 n-1，由 1 退化的线性链表可以反推。
含有 n≥1 个结点的二叉树的高度至少为 logn，由 2 满二叉树可以反推。
在二叉树的第 i 层，至多有 2^(i-1) 个节点，比如该层是满的。

对于一棵有 n 个节点的完全二叉树，从上到下，从左到右进行序号编号，
对于任一个节点，编号 i=0 表示树根节点，
编号 i 的节点的左右儿子节点编号分别为：2i+1,2i+2，
父亲节点编号为：i/2，整除操作去掉小数
*/

// TreeNode 二叉树可以使用链表来实现
// 数组也可以用来表示二叉树，一般用来表示完全二叉树
// 二叉树结构体
type TreeNode struct {
	Data string // 结点用来存放数据
	Left *TreeNode // 左子树
	Right *TreeNode // 右子树
}

/*一般使用二叉树来实现查找的功能
遍历二叉树
构建一棵树后，我们希望遍历它，有四种遍历方法：
先序遍历：先访问根节点，再访问左子树，最后访问右子树。
后序遍历：先访问左子树，再访问右子树，最后访问根节点。
中序遍历：先访问左子树，再访问根节点，最后访问右子树。
层次遍历：每一层从左到右访问每一个节点
*/

// PreOrder 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 先打印根节点
	fmt.Printf("%s-", tree.Data)
	// 在打印左子树
	PreOrder(tree.Left)
	// 最后打印右子树
	PreOrder(tree.Right)
}

// MidOrder 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 先打印左子树
	MidOrder(tree.Left)
	// 在打印根节点
	fmt.Printf("%s-", tree.Data)
	// 最后打印右子树
	MidOrder(tree.Right)
}

// PostOrder 后续遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 先打印左子树
	PreOrder(tree.Left)
	// 在打印右子树
	PostOrder(tree.Right)
	// 最后打印根节点
	fmt.Printf("%s-", tree.Data)
}

// 测试
func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	MidOrder(t)
	fmt.Println("\n后序排序")
	PostOrder(t)
}

















