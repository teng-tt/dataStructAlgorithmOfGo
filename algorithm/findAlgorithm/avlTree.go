package main

import (
	"fmt"
)

/*
AVL树
二叉查找树的树高度影响了查找的效率，需要尽量减小树的高度，AVL树正是这样的树

AVL树是一棵严格自平衡的二叉查找树：
1、首先它是一棵二叉查找树。
2、任意一个节点的左右子树最大高度差为1。

由于树特征定义，可以计算出其高度 h 的上界 h<=1.44log(n)，也就是最坏情况下，树的高度约等于 1.44log(n)
树的高度被限制于 1.44log(n)，查找元素时使用二分查找，最坏查找 1.44log(n) 次，最坏时间复杂度为 1.44log(n)，
去掉常数项，时间复杂度为：log(n)。

为了维持AVL树的特征，每次添加和删除元素都需要一次或多次旋转来调整树的平衡
调整的依据来自于二叉树节点的平衡因子：节点的左子树与右子树的高度差称为该节点的平衡因子
约束范围为 [-1，0，1]

平衡二叉查找树比较难以理解的是添加和删除元素时的调整操作
 */

// AVL树的基本数据结构

// AVLTree AVL树
type AVLTree struct {
	Root *AVLTreeNode // 树的根节点
}

// AVLTreeNode AVL节点
type AVLTreeNode struct {
	Value int64			// 值
	Times int64			// 值出现的次数
	Height int64		// 该节点作为树根节点，树的高度，方便计算平衡因子
	Left *AVLTreeNode	// 左子树
	Right *AVLTreeNode	// 右子树
}

// NewAVLTree 初始化一个AVL树
func NewAVLTree() *AVLTree {
	return new(AVLTree)
}

// UpdateHeight 更新树的高度
func (node *AVLTreeNode) UpdateHeight() {
	if node == nil {
		return
	}
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	// 那个子树高，算哪棵的
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	// 高度加上自己那一层
	node.Height = maxHeight + 1
}

// BalanceFactor 计算树的平衡因子，也就是左右子树的高度差
func (node *AVLTreeNode) BalanceFactor() int64 {
	var leftHeight, rightHeight int64 = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}


/*
AVL树添加元素
添加元素前需要定位到元素的位置，也就是使用二分查找找到该元素需要插入的地方。
插入后，需要满足所有节点的平衡因子在 [-1，0，1] 范围内，如果不在，需要进行旋转调整。
旋转有四种情况：
1、在右子树上插上右儿子导致失衡，左旋，转一次。
2、在左子树上插上左儿子导致失衡，右旋，转一次。
3、在左子树上插上右儿子导致失衡，先左后右旋，转两次。
4、在右子树上插上左儿子导致失衡，先右后左旋，转两次。

#旋转规律记忆法：单旋和双旋，单旋反方向，双旋同方向
 */

// RightRotation 左左情况: 左子树插左儿子：单右旋
// 将 Pivot 代替 旧root 的位置成为新的 Root
// 旧 root 委屈一下成为 Pivot 的右儿子
// 而 Pivot 的右儿子变成了 原来 root 的左儿子
// 相应调整后树的高度降低了，该失衡消失
func RightRotation(root *AVLTreeNode) *AVLTreeNode {
	// 只有Pivot和B，Root位置变了
	Pivot := root.Left
	B := Pivot.Right
	Pivot.Right = root
	root.Left = B
	// 只有Root和Pivot变化了高度
	root.UpdateHeight()
	Pivot.UpdateHeight()

	return Pivot
}

// LeftRotation 右右情况：右子树插右儿子：单左旋
// 将 Pivot 代替 旧root 的位置成为新的 Root
// 旧 root 委屈一下成为 Pivot 的左儿子
// 而 Pivot 的左儿子变成了 原来 root 的右儿子
// 相应调整后树的高度降低了，该失衡消失
func LeftRotation(root *AVLTreeNode) *AVLTreeNode {
	// 只有Pivot和B，Root位置变了
	Pivot := root.Right
	B := Pivot.Left
	Pivot.Left = root
	root.Right = B
	// 只有Root和Pivot变化了高度
	root.UpdateHeight()
	Pivot.UpdateHeight()

	return Pivot
}

// LeftRightRotation 左右情况：左子树插右儿子：先左后右旋
// 直接复用了之前左旋和右旋的代码，虽然难以理解，但是画一下图
// 确实这样调整后树高度降了，不再失衡
func LeftRightRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

// RightLeftRotation 右左情况：右子树插左儿子：先右后左旋
func RightLeftRotation(node *AVLTreeNode) *AVLTreeNode {
	node.Right = RightRotation(node.Right)
	return LeftRotation(node)
}

// Add 四种旋转代码实现后，我们开始进行添加元素操作
// 添加元素
func (tree *AVLTree) Add(value int64) {
	// 往树根添加元素，会返回新的树根
	tree.Root = tree.Root.Add(value)
}

func (node *AVLTreeNode) Add(value int64) *AVLTreeNode {
	// 添加值到根节点node，如果node为空，那么让值成为新的根节点，树的高度为1
	if node == nil {
		return &AVLTreeNode{Value: value, Height: 1}
	}
	// 如果值重复，什么都不用做，直接更新次数
	if node.Value == value {
		node.Times = node.Times + 1
		return node
	}
	// 辅助变量
	var newTreeNode *AVLTreeNode
	if value > node.Value {
		// 插入的值大于节点值，要从右子树继续插入
		node.Right = node.Right.Add(value)
		// 平衡因子，插入右子树后，要确保树根左子树的高度不能比右子树低一层
		factor := node.BalanceFactor()
		// 右子树的高度变高了，导致左子树-右子树的高度从-1变成了-2
		if factor == -2 {
			if value > node.Right.Value {
				// 表示在右子树上插上右儿子导致失衡，需要单左旋：
				newTreeNode = LeftRotation(node)
			}else {
				//表示在右子树上插上左儿子导致失衡，先右后左旋：
				newTreeNode = RightLeftRotation(node)
			}
		}
	}else {
		// 插入的值小于节点值，要从左子树继续插入
		node.Left = node.Left.Add(value)
		// 平衡因子，插入左子树后，要确保树根左子树的高度不能比右子树高一层。
		factor := node.BalanceFactor()
		// 左子树的高度变高了，导致左子树-右子树的高度从1变成了2。
		if factor == 2 {
			if value < node.Left.Value {
				// 表示在左子树上插上左儿子导致失衡，需要单右旋：
				newTreeNode = RightRotation(node)
			}else {
				//表示在左子树上插上右儿子导致失衡，先左后右旋：
				newTreeNode = LeftRightRotation(node)
			}
		}
	}
	if newTreeNode == nil {
		// 表示什么旋转都没有，根节点没变，直接刷新树高度
		node.UpdateHeight()
		return node
	}else {
		// 旋转了，树根节点变了，需要刷新新的树根高度
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

// AVL树查找元素等操作
// 查找操作逻辑与通用的二叉查找树一样，并无区别

// FindMinValue 找出最小值的结点
func (tree *AVLTree) FindMinValue() *AVLTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *AVLTreeNode) FindMinValue() *AVLTreeNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}
	// 一直左子树递归
	return node.Left.FindMinValue()
}

// FindMaxValue 找出最大值节点
func (tree *AVLTree) FindMaxValue() *AVLTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *AVLTreeNode) FindMaxValue() *AVLTreeNode {
	if node.Right == nil {
		// 右子树为空，表面已经是最右的节点了，该值就是最大值
		return node
	}
	// 一直右子树递归
	return node.Right.FindMaxValue()
}

// Find 查找指定节点
func (tree *AVLTree) Find(value int64) *AVLTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return  nil
	}
	return tree.Root.Find(value)
}

func (node *AVLTreeNode) Find(value int64) *AVLTreeNode {
	if node.Value == value {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node
	} else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		// 递归左子树查找
		return node.Left.Find(value)
	}else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if value > node.Value {
			if node.Right == nil {
				// 右子树为空，表示找不到该值了，返回nil
				return  nil
			}
		}
		// 递归右子树查找
		return node.Right.Find(value)
	}
}

// MidOrder 中序遍历，可以实现排序
func (tree *AVLTree) MidOrder() {
	tree.Root.MidOreder()
}

func (node *AVLTreeNode) MidOreder() {
	if node == nil {
		return
	}
	// 先打印左子树
	node.Left.MidOreder()
	// 按照次数打印根节点
	for i := 0; i < int(node.Times); i++ {
		fmt.Println("value:", node.Value, " tree height:", node.BalanceFactor())
	}
	// 最后打印右子树
	node.Right.MidOreder()
}

/*
AVL树删除操作
删除元素有四种情况：
1、删除的节点是叶子节点，没有儿子，直接删除后看离它最近的父亲节点是否失衡，做调整操作。
2、删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，如果左子树更高，
选择左子树中最大的节点，也就是左子树最右边的叶子节点，如果右子树更高，选择右子树中最小的节点，
也就是右子树最左边的叶子节点。最后，删除这个叶子节点，也就是变成情况1。

3、删除的节点只有左子树，可以知道左子树其实就只有一个节点，被删除节点本身
（假设左子树多于2个节点，那么高度差就等于2了，不符合AVL树定义），将左节点
替换被删除的节点，最后删除这个左节点，变成情况1。

4、删除的节点只有右子树，可以知道右子树其实就只有一个节点，被删除节点本身
（假设右子树多于2个节点，那么高度差就等于2了，不符合AVL树定义），将右节点
替换被删除的节点，最后删除这个右节点，变成情况1。

后面三种情况最后都变成 情况1，就是将删除的节点变成叶子节点，然后可以直接删除该叶子节点，
然后看其最近的父亲节点是否失衡，失衡时对树进行平衡
 */


/*
当删除的值不等于当前节点的值时，在相应的子树中递归删除，递归过程中会自底向上维护AVL树的特征。

1、小于删除的值 value < node.Value，在左子树中递归删除：node.Left = node.Left.Delete(value)。
2、大于删除的值 value > node.Value，在右子树中递归删除：node.Right = node.Right.Delete(value)。
因为删除后可能因为旋转调整，导致树根节点变了，这时会返回新的树根，递归删除后需要将返回的新根节点赋予原来的老根节点。

情况1，找到要删除的值时，该值是叶子节点，直接删除该节点即可：
情况2，删除的节点有两棵子树，选择高度更高的子树下的节点来替换被删除的节点：
情况3和情况4，如果被删除的节点只有一个子树，那么该子树一定没有儿子，
不然树的高度就大于1了，所以直接替换值后删除该子树节点：

核心在于删除后的旋转调整，如果删除的值不匹配当前节点的值，
对当前节点的左右子树进行递归删除，递归删除后该节点为根节点的子树可能不平衡，
需要判断后决定要不要旋转这棵树。
每次递归都是自底向上，从很小的子树到很大的子树，如果自底向上每棵子树都进行调整，
约束在树的高度差不超过1，那么整棵树自然也符合AVL树的平衡规则。

删除元素后，如果子树失衡，需要进行调整操作，主要有两种：删除后左子树比右子树高，删除后右子树比左子树高。
 */

func (tree *AVLTree) Delete(value int64)  {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}
	tree.Root = tree.Root.Delete(value)
}

// Delete 删除元素
func (node *AVLTreeNode) Delete(value int64) *AVLTreeNode {
	if node == nil {
		// 如果是空树，直接返回
		return nil
	}
	if value < node.Value {
		// 从左子树开始删除
		node.Left = node.Left.Delete(value)
		// 删除后要更新该子树高度
		node.Left.UpdateHeight()
	}else if value > node.Value {
		// 从右子树开始删除
		node.Right = node.Right.Delete(value)
		// 删除后要更新该子树高度
		node.Right.UpdateHeight()
	}else {
		// 找到该值对应的节点
		// 该节点没有左右子树
		// 第一种情况，删除的节点没有儿子，直接删除即可
		if node.Left == nil && node.Right == nil {
			return nil // 直接返回nil，表示直接该值删除
		}

		// 该节点有两棵子树，选择更高的哪个来替换
		// 第二种情况，删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，
		// 如果左子树更高，选择左子树中最大的节点，也就是左子树最右边的叶子节点，
		// 如果右子树更高，选择右子树中最小的节点，也就是右子树最左边的叶子节点。
		// 最后，删除这个叶子节点。
		if node.Left != nil && node.Right != nil {
			// 左子树更高，拿左子树中最大值的节点替换
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}
				// 最大值的节点替换被删除节点
				node.Value = maxNode.Value
				node.Times = maxNode.Times
				// 把最大值的节点删掉
				node.Left = node.Left.Delete(maxNode.Value)
				// 删除后要更新该子树高度
				node.Left.UpdateHeight()
			}else {
				// 右子树更高，拿右子树中最小值的节点替换
				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}
				// 最小值的节点替换被删除节点
				node.Value = minNode.Value
				node.Times = minNode.Times
				// 把最小的节点删掉
				node.Right = node.Right.Delete(minNode.Value)
				// 删除后要更新该子树高度
				node.Right.UpdateHeight()
			}
		}else {
			// 只有左子树或只有右子树
			// 只有一个子树，该子树也只是一个节点，将该节点替换被删除的节点，然后置子树为空
			if node.Left != nil {
				// 第三种情况，删除的节点只有左子树，因为树的特征，
				// 可以知道左子树其实就只有一个节点，它本身，否则高度差就等于2了
				node.Value = node.Left.Value
				node.Times = node.Left.Times
				node.Height = 1
				node.Left = nil
			}else if node.Right != nil {
				// 第四种情况，删除的节点只有右子树，因为树的特征，
				// 可以知道右子树其实就只有一个节点，它本身，否则高度差就等于2了
				node.Value = node.Right.Value
				node.Times = node.Right.Times
				node.Height = 1
				node.Right = nil
			}
		}
		// 找到值后，进行替换删除后，直接返回该节点
		return node
	}
	// 左右子树递归删除节点后需要平衡
	var newNode *AVLTreeNode
	// 相当删除了右子树的节点，左边比右边高了，不平衡
	if node.BalanceFactor() == 2 {
		if node.Left.BalanceFactor() >= 0 {
			newNode = RightRotation(node)
		}else {
			newNode = LeftRightRotation(node)
		}
		//  相当删除了左子树的节点，右边比左边高了，不平衡
	}else if node.BalanceFactor() == -2 {
		if node.Right.BalanceFactor() <= 0 {
			newNode = LeftRotation(node)
		}else {
			newNode = RightLeftRotation(node)
		}
	}
	if newNode == nil {
		node.UpdateHeight()
		return node
	}else {
		newNode.UpdateHeight()
		return newNode
	}
}
/*
时间复杂度分析
删除操作是先找到删除的节点，然后将该节点与一个叶子节点交换，接着删除叶子节点，
最后对叶子节点的父层逐层向上旋转调整。

删除操作的时间复杂度和添加操作一样。区别在于，添加操作最多旋转两次就可以达到树的平衡，
而删除操作可能会旋转超过两次。
 */

// IsAVLTree 验证是否是一棵AVL树
func (tree *AVLTree) IsAVLTree() bool {
	if tree == nil || tree.Root == nil {
		return true
	}
	// 判断节点是否符合 AVL 树的定义
	if tree.Root.IsRight() {
		return true
	}
	return false
}

// IsRight 判断节点是否符合 AVL 树的定义
func (node *AVLTreeNode) IsRight() bool {
	if node == nil {
		return true
	}
	// 左右子树都为空，那么是叶子节点
	if node.Left == nil && node.Right == nil {
		// 叶子节点高度应该为1
		if node.Height == 1{
			return true
		}else {
			fmt.Println("leaf node height is ", node.Height)
			return false
		}
	}else if node.Left != nil && node.Right != nil {
		// 左右子树都是满的
		// 左儿子必须比父亲小，右儿子必须比父亲大
		if node.Left.Value < node.Value && node.Right.Value > node.Value {
		}else {
			// 不符合 AVL 树定义
			fmt.Printf("father is %v lchild is %v, rchild is %v\n", node.Value, node.Left.Value, node.Right.Value)
			return false
		}
		bal := node.Left.Height - node.Right.Height
		if bal < 0 {
			bal = -bal
		}
		// 子树高度差不能大于1
		if bal > 1 {
			fmt.Println("sub tree height bal is ", bal)
			return false
		}
		// 如果左子树比右子树高，那么父亲的高度等于左子树+1
		if node.Left.Height > node.Right.Height {
			if node.Height == node.Left.Height + 1 {
			}else {
				fmt.Printf("%#v height:%v,left sub tree height: %v,right sub tree height:%v\n", node, node.Height, node.Left.Height, node.Right.Height)
				return false
			}
		} else {
			// 如果右子树比左子树高，那么父亲的高度等于右子树+1
			if node.Height == node.Right.Height + 1 {
			}else {
				fmt.Printf("%#v height:%v,left sub tree height: %v,right sub tree height:%v\n", node, node.Height, node.Left.Height, node.Right.Height)
				return false
			}
		}
		// 递归判断子树
		if !node.Left.IsRight() {
			return false
		}
		// 递归判断子树
		if !node.Right.IsRight() {
			return false
		}
	}else {
		// 只存在一棵子树
		if node.Right != nil {
			// 子树高度只能是1
			if node.Right.Height == 1 && node.Right.Left == nil && node.Right.Right == nil {
				if node.Right.Value > node.Value {
					// 右节点必须比父亲大
				}else {
					fmt.Printf("%v,(%#v,%#v) child", node.Value, node.Right, node.Left)
					return false
				}
			}else {
				fmt.Printf("%v,(%#v,%#v) child", node.Value, node.Right, node.Left)
				return false
			}
		}else {
			if node.Left.Height == 1 && node.Left.Left == nil && node.Left.Right == nil {
				if node.Left.Value < node.Value {
					// 左节点必须比父亲小
				}else {
					fmt.Printf("%v,(%#v,%#v) child", node.Value, node.Right, node.Left)
					return false
				}
			}else {
				fmt.Printf("%v,(%#v,%#v) child", node.Value, node.Right, node.Left)
				return false
			}
		}
	}
	return true
}

// 验证测试
// 程序是递归程序，如果改写为非递归形式，效率和性能会更好，
// 在此就不实现了，理解AVL树添加和删除的总体思路即可
// AVL 树作为严格平衡的二叉查找树，在 windows 对进程地址空间的管理被使用到
func main() {
	values := []int64{2, 3, 7, 10, 10, 10, 10, 23, 9, 102, 109, 111, 112, 113}

	// 初始化二叉查找树并添加元素
	tree := NewAVLTree()
	for _, v := range values {
		tree.Add(v)
	}

	// 找到最大值或最小值的节点
	fmt.Println("find min value:", tree.FindMinValue())
	fmt.Println("find max value:", tree.FindMaxValue())

	// 查找不存在的99
	node := tree.Find(99)
	if node != nil {
		fmt.Println("find it 99!")
	} else {
		fmt.Println("not find it 99!")
	}

	// 查找存在的9
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	// 删除存在的9后，再查找9
	tree.Delete(9)
	tree.Delete(10)
	tree.Delete(2)
	tree.Delete(3)
	tree.Add(4)
	tree.Add(3)
	tree.Add(10)
	tree.Delete(111)
	node = tree.Find(9)
	if node != nil {
		fmt.Println("find it 9!")
	} else {
		fmt.Println("not find it 9!")
	}

	// 中序遍历，实现排序
	tree.MidOrder()

	if tree.IsAVLTree() {
		fmt.Println("is a avl tree")
	} else {
		fmt.Println("is not avl tree")
	}
}