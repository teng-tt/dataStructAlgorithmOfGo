package main

import (
	"fmt"
)

/*
左倾红黑树
左倾红黑树可以由 2-3 树的二叉树形式来实现。

其定义为：
1、根节点的链接是黑色的
2、红链接均为左链接
3、没有任何一个结点同时和两条红链接相连
4、任意一个节点到达叶子节点的所有路径，经过的黑链接数量相同，也就是该树是完美黑色平衡的
比如，某一个节点，它可以到达5个叶子节点，那么这5条路径上的黑链接数量一样。

5、由于红链接都在左边，所以这种红黑树又称左倾红黑树。
左倾红黑树与 2-3 树一一对应，只要将左链接画平
 */

// 使用2-3树的形式实现左倾红黑树
// 具体实现如下

// 节点旋转和颜色转换
// 首先，我们要定义树的结构 LLRBTree ，
// 以及表示左倾红黑树的节点 LLRBTNode

// 定义颜色
const (
	RED = true
	BLACK = false
)

// LLRBTree 左倾红黑树
type LLRBTree struct {
	Root *LLRBTNode // 树的根节点
}

// LLRBTNode 左倾红黑树节点
type LLRBTNode struct {
	Value int64		// 值
	Times int64		// 值出现的次数
	Left *LLRBTNode	// 左子树
	Right *LLRBTNode// 右子树
	Color bool		// 父亲指向该节点的链接颜色
}

// NewLLRBTree 新建一棵树
func NewLLRBTree() *LLRBTree {
	return &LLRBTree{}
}

// IsRed 节点颜色
func IsRed(node *LLRBTNode) bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}
/*
在节点 LLRBTNode 中，我们存储的元素字段为 Value，
由于可能有重复的元素插入，所以多了一个 Times 字段，表示该元素出现几次。

当然，红黑树中的红黑颜色使用 Color 定义，表示父亲指向该节点的链接颜色。
为了方便，我们还构造了一个辅助函数 IsRed()。

在元素添加和实现的过程中，需要做调整操作，有两种旋转操作，
对某节点的右链接进行左旋转，或者左链接进行右旋转

旋转和颜色转换作为局部调整，并不影响全局
 */

// RotationLeft 左旋转
func RotationLeft(h *LLRBTNode) *LLRBTNode {
	if h == nil {
		return nil
	}
	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = h.Color
	h.Color = RED
	return x
}

// RotationRight 右旋转
func RotationRight(h *LLRBTNode) *LLRBTNode {
	if h == nil {
		return nil
	}
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = h.Color
	h.Color = RED
	return x
}

// ColorChange 颜色转换
// 由于左倾红黑树不允许一个节点有两个红链接，所以需要做颜色转换
func ColorChange(h *LLRBTNode) {
	if h == nil {
		return
	}
	h.Color = !h.Color
	h.Left.Color = !h.Left.Color
	h.Right.Color = !h.Right.Color
}


/*
添加元素
每次添加元素节点时，都将该节点 Color 字段，也就是父亲指向它的链接设置为 RED 红色。
接着判断其父亲是否有两个红链接（如连续的两个左红链接或者左右红色链接），
或者有右红色链接，进行颜色变换或旋转操作。

主要有以下这几种情况。
插入元素到2节点，
1、向左插入，直接让节点变为3节点，直接插入
2、向右插入，直接让节点变为3节点，直接插入，不过右插入时需要左旋使得红色链接在左边
插入元素到3节点
需要做旋转和颜色转换操作
1、新键最大，插入到3节点最右边，用红链接与新节点相连接，最后将链接颜色变为黑色
2、新键最小，插入到3节点的最左边，用红链接与新节点相连接，
3节点进行右旋转变为红色右链接，最后将链接颜色变为黑色

3、新键介于两者之间，插入到3节点的最左子树的右叶子节点，用红链接与新节点相连接，
3节点进行左旋转变为红色左链接，然后进行右旋转变为红色右链接
最后将链接颜色变为黑色
 */

// Add 左倾红黑树添加元素
func (tree *LLRBTree) Add(value int64) {
	// 根节点开始添加元素，因为可能调整，所以需要将返回的节点赋值回根节点
	tree.Root = tree.Root.Add(value)
	// 根节点的链接永远都是黑色的
	tree.Root.Color = BLACK
}

// Add 往节点添加元素
func (node *LLRBTNode) Add(value int64) *LLRBTNode {
	// 插入的节点为空，将其链接颜色设置为红色，并返回
	if node == nil {
		return &LLRBTNode{
			Value:value,
			Color:RED,
		}
	}
	// 插入的元素重复
	if value == node.Value {
		node.Times = node.Times + 1
	}else if value > node.Value {
		// 插入的元素比节点值大，往右子树插入
		node.Right = node.Right.Add(value)
	}else {
		// 插入的元素比节点值小，往左子树插入
		node.Left = node.Left.Add(value)
	}
	// 辅助变量
	nowNode := node
	// 右链接为红色，那么进行左旋，确保树是左倾的
	// 这里做完操作后就可以结束了，因为插入操作，新插入的右红链接左旋后，
	// nowNode节点不会出现连续两个红左链接，因为它只有一个左红链接
	if IsRed(nowNode.Right) && !IsRed(node.Left) {
		nowNode = RotationLeft(nowNode)
	}else {
		// 连续两个左链接为红色，那么进行右旋
		if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
			nowNode = RotationRight(nowNode)
		}
		// 旋转后，可能左右链接都为红色，需要变色
		if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
			ColorChange(nowNode)
		}
	}
	return nowNode
}
/*
左倾红黑树的最坏树高度为 2log(n)，其中 n 为树的节点数量。为什么呢，我们先把左倾红黑树当作 2-3 树，
也就是说最坏情况下沿着 2-3 树左边的节点都是3节点，其他节点都是2节点，这时树高近似 log(n)，
再从 2-3 树转成左倾红黑树，当3节点不画平时，可以知道树高变成原来 2-3 树树高的两倍。
虽然如此，构造一棵最坏的左倾红黑树很难

我们的代码实现中，左倾红黑树的插入，需要逐层判断是否需要旋转和变色，复杂度为 log(n)
对于 AVL 树来说，插入最多旋转两次，但其需要逐层更新树高度，复杂度也是为 log(n)。
我们不再纠结两种平衡树哪种更好，因为代码实现中，两种平衡树都需要自底向上的递归操作，效率差别不大
*/

/*
删除元素
删除操作就复杂得多了。对照一下 2-3 树
情况1：如果删除的是非叶子节点，找到其最小后驱节点，也就是在其右子树中一直向左找，
找到的该叶子节点替换被删除的节点，然后删除该叶子节点，变成情况2。

情况2：如果删除的是叶子节点，如果它是红节点，也就是父亲指向它的链接为红色，那么直接删除即可。
否则，我们需要进行调整，使它变为红节点，再删除。

在这里，为了使得删除叶子节点时可以直接删除，叶子节点必须变为红节点
（在 2-3 树中，也就是2节点要变成3节点，我们知道要不和父亲合并再递归向上，要不向兄弟借值然后重新分布）

我们创造两种操作，
如果删除的节点在左子树中，可能需要进行红色左移，
如果删除的节点在右子树中，可能需要进行红色右移。

红色左移的步骤：
要在树 h 的的左子树中删除元素，这时树 h 根节点是红节点，其儿子 b，d 节点都为黑色节点，

且两个黑色节点都是2节点，都没有左红孩子，那么直接对 h 树根节点变色即可
（相当于 2-3 树：把父亲的一个值拉下来合并）

如果存在右儿子 d 是3节点，有左红孩子 e，那么需要先对 h 树根节点变色后，
对右儿子 d 右旋，再对 h 树根节点左旋，最后再一次对 h 树根节点变色
（相当于 2-3 树：向3节点兄弟借值，然后重新分布）

我们知道 2-3 树的删除是从叶子节点开始，自底向上的向兄弟节点借值，或和父亲合并，然后一直递归到根节点。
左倾红黑树参考了这种做法，但更巧妙，
左倾红黑树要保证一路上每次递归进入删除操作的子树树根一定是一个3节点，
所以需要适当的红色左移或右移（类似于 2-3 树借值和合并），
这样一直递归到叶子节点，叶子节点也会是一个3节点，然后就可以直接删除叶子节点，
最后再自底向上的恢复左倾红黑树的特征。

树h 左子树和右子树删除元素分别有四种情况，后两种情况需要使用到红色左移或右移，
状态演变之后， 树h 才可以从左或右子树进入下一次递归
 */

// MoveRedLeft 红色左移
// 节点 h 是红节点，其左儿子和左儿子的左儿子都为黑节点，
// 左移后使得其左儿子或左儿子的左儿子有一个是红色节点
// 为什么要红色左移，是要保证调整后，子树根节点 h 的左儿子
// 或左儿子的左儿子有一个是红色节点，这样从 h 的左子树递归删除元素才可以继续下去
func MoveRedLeft(h *LLRBTNode) *LLRBTNode {
	// 应该确保 isRed(h) && !isRed(h.left) && !isRed(h.left.left)
	ColorChange(h)
	// 右儿子有左红链接
	if IsRed(h.Right.Left) {
		// 对右儿子右旋
		h.Right = RotationRight(h.Right)
		// 再左旋
		h = RotationLeft(h)
	}
	return h
}

// MoveRedRight 红色右移
// 节点 h 是红节点，其右儿子和右儿子的左儿子都为黑节点，
// 右移后使得其右儿子或右儿子的右儿子有一个是红色节点
// 为什么要红色右移，同样是为了保证树根节点 h 的右儿子
// 或右儿子的右儿子有一个是红色节点，往右子树递归删除元素可以继续下去
func MoveRedRight(h *LLRBTNode) *LLRBTNode {
	// 应该确保 isRed(h) && !isRed(h.right) && !isRed(h.right.left)
	ColorChange(h)
	// 左儿子有左红链接
	if IsRed(h.Left.Left) {
		// 右旋
		h = RotationRight(h)
		// 变色
		ColorChange(h)
	}
	return h
}

// Delete 左倾红黑树删除元素
func (tree *LLRBTree) Delete(value int64) {
	// 当找不到值时直接返回
	if tree.Find(value) == nil {
		return
	}
	if !IsRed(tree.Root.Left) && !IsRed(tree.Root.Right) {
		// 左右子树都是黑节点，那么先将根节点变为红节点，方便后面的红色左移或右移
		tree.Root.Color = RED
	}
	tree.Root = tree.Root.Delete(value)
	// 最后，如果根节点非空，永远都要为黑节点，赋值黑色
	if tree.Root != nil {
		tree.Root.Color = BLACK
	}
}

// 首先 tree.Find(value) 找到可以删除的值时才能进行删除。
// 当根节点的左右子树都为黑节点时，那么先将根节点变为红节点，方便后面的红色左移或右移。
// 删除完节点：tree.Root = tree.Root.Delete(value) 后，需要将根节点染回黑色，
// 因为左倾红黑树的特征之一是根节点永远都是黑色。

// Delete 核心的从子树中删除元素代码如下
func (node *LLRBTNode) Delete(value int64) *LLRBTNode {
	// 辅助变量
	nowNode := node
	// 删除的元素比子树根节点小，需要从左子树删除
	if value < node.Value {
		// 因为从左子树删除，所以要判断是否需要红色左移
		if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
			// 左儿子和左儿子的左儿子都不是红色节点，那么没法递归下去，先红色左移
			nowNode = MoveRedLeft(nowNode)
		}
		// 现在可以从左子树中删除了
		nowNode.Left = nowNode.Left.Delete(value)
	}else {
		// 删除的元素等于或大于树根节点
		// 左节点为红色，那么需要右旋，方便后面可以红色右移
		if IsRed(nowNode.Left) {
			nowNode = RotationRight(nowNode)
		}
		// 值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
		// 为什么呢，反证，它没有右儿子，但有左儿子，因为左倾红黑树的特征，
		// 那么左儿子一定是红色，但是前面的语句已经把红色左儿子右旋到右边，不应该出现右儿子为空
		if value == nowNode.Value && nowNode.Right == nil {
			return nil
		}
		// 因为从右子树删除，所以要判断是否需要红色右移
		if !IsRed(nowNode.Right) && !IsRed(nowNode.Right.Left) {
			// 右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
			nowNode = MoveRedRight(nowNode)
		}
		// 删除的节点找到了，它是中间节点，需要用最小后驱节点来替换它，然后删除最小后驱节点
		if value == nowNode.Value {
			minNode := nowNode.Right.FindMinValue()
			nowNode.Value = minNode.Value
			nowNode.Times = minNode.Times
			// 删除其最小后驱节点
			nowNode.Right = nowNode.Right.DeleteMin()
		}else {
			// 删除的元素比子树根节点大，需要从右子树删除
			// 如果不是删除内部节点，依然是从右子树继续递归
			nowNode.Right = nowNode.Right.Delete(value)
		}
	}
	// 递归完成后还要进行一次 FixUp()，恢复左倾红黑树的特征
	// 最后，删除叶子节点后，恢复左倾红黑树特征
	return nowNode.FixUp()
}

// DeleteMin 删除最小节点
func (node *LLRBTNode) DeleteMin() *LLRBTNode {
	// 辅助变量
	nowNode := node
	// 没有左子树，那么删除自己
	if nowNode.Left == nil {
		return nil
	}
	// 判断是否需要红色左移，因为最小元素在左子树中
	if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
		nowNode = MoveRedLeft(nowNode)
	}
	// 递归从左子树删除
	// 因为最小节点在最左的叶子节点，所以只需要适当的红色左移，然后一直左子树递归即可
	nowNode.Left = nowNode.Left.DeleteMin()
	// 修复左倾红黑树特征
	return nowNode.FixUp()
}

// FixUp 修复左倾红黑树特征
func (node *LLRBTNode) FixUp() *LLRBTNode {
	// 辅助变量
	nowNode := node
	// 红链接在右边，左旋恢复，让红链接只出现在左边
	if IsRed(nowNode.Right) {
		nowNode = RotationLeft(nowNode)
	}
	// 连续两个左链接为红色，那么进行右旋
	if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
		nowNode = RotationRight(nowNode)
	}
	// 旋转后，可能左右链接都为红色，需要变色
	if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
		ColorChange(nowNode)
	}
	return nowNode
}
/*
删除操作很难理解，可以多多思考，红色左移和右移不断地递归都是为了确保删除叶子节点时，其是一个3节点。
PS：如果不理解自顶向下的红色左移和右移递归思路，可以更换另外一种方法，使用原先 2-3树 删除元素操作
步骤来实现，一开始从叶子节点删除，然后自底向上的向兄弟借值或与父亲合并，这是更容易理解的，我们不在
这里进行展示了，可以借鉴普通红黑树章节的删除实现（它使用了自底向上的调整）

左倾红黑树删除元素需要自顶向下的递归，可能不断地红色左移和右移，也就是有很多的旋转，
当删除叶子节点后，还需要逐层恢复左倾红黑树的特征。时间复杂度仍然是和树高有关：log(n)
 */


// 查找元素等实现，查找最小值，最大值，或者某个值
// 查找操作逻辑与通用的二叉查找树一样，并无区别。

// FindMinValue 找出最小值的节点
func (tree *LLRBTree) FindMinValue() *LLRBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMinValue()
}

func (node *LLRBTNode) FindMinValue() *LLRBTNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}
	// 一直左子树递归
	return node.Left.FindMinValue()
}

// FindMaxValue 找出最大值的节点
func (tree *LLRBTree) FindMaxValue() *LLRBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMaxValue()
}

func (node *LLRBTNode) FindMaxValue() *LLRBTNode {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}
	// 一直右子树递归
	return node.Right.FindMaxValue()
}

// Find 查找指定节点
func (tree *LLRBTree) Find(value int64) *LLRBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.Find(value)
}

func (node *LLRBTNode) Find(value int64) *LLRBTNode {
	if value == node.Value {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node
	}else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Left.Find(value)
	}else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return  nil
		}
		return node.Right.Find(value)
	}
}

// MidOrder 中序遍历
func (tree *LLRBTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *LLRBTNode) MidOrder() {
	if node == nil {
		return
	}
	// 先打印左子树
	node.Left.MidOrder()
	// 按次数打印根节点
	for i := 0; i < int(node.Times); i++ {
		fmt.Println(node.Value)
	}
	// 打印右子树
	node.Right.MidOrder()
}

// IsLLRBTree 验证是否是一棵左倾红黑树
func (tree *LLRBTree) IsLLRBTree() bool {
	if tree == nil || tree.Root == nil {
		return true
	}
	// 判断树是否是一棵二分查找树
	if !tree.Root.IsBST() {
		return false
	}
	// 判断树是否遵循2-3树，也就是红链接只能在左边，不能连续有两个红链接
	if !tree.Root.Is23() {
		return false
	}
	// 判断树是否平衡，也就是任意一个节点到叶子节点，经过的黑色链接数量相同
	// 先计算根节点到最左边叶子节点的黑链接数量
	balckNum := 0
	x := tree.Root
	for x != nil {
		if !IsRed(x) {
			// 是黑色链接
			balckNum = balckNum + 1
		}
		x = x.Left
	}
	if !tree.Root.IsBalanced(balckNum) {
		return false
	}
	return true
}

// IsBST 节点所在的子树是否是一棵二分查找树
func (node *LLRBTNode) IsBST() bool {
	if node == nil {
		return true
	}
	// 左子树非空，那么根节点必须大于左儿子节点
	if node.Left != nil {
		if node.Value > node.Left.Value{
		}else {
			fmt.Printf("father:%#v,lchild:%#v,rchild:%#v\n", node, node.Left, node.Right)
			return false
		}
	}
	// 右子树非空，那么根节点必须小于右儿子节点
	if node.Right != nil {
		if node.Value < node.Right.Value {
		}else {
			fmt.Printf("father:%#v,lchild:%#v,rchild:%#v\n", node, node.Left, node.Right)
			return false
		}
	}
	// 左子树也要判断是否是平衡查找树
	if !node.Left.IsBST() {
		return false
	}
	// 右子树也要判断是否是平衡查找树
	if !node.Right.IsBST() {
		return false
	}
	return true
}

// Is23 节点所在的子树是否遵循2-3树
func (node *LLRBTNode) Is23() bool {
	if node == nil {
		return true
	}
	// 不允许右倾红链接
	if IsRed(node.Right) {
		fmt.Printf("father:%#v,rchild:%#v\n", node, node.Right)
		return false
	}
	// 不允许连续两个左红链接
	if IsRed(node.Left) && IsRed(node.Left.Left) {
		fmt.Printf("father:%#v,lchild:%#v\n", node, node.Left)
		return false
	}
	// 左子树也要判断是否遵循2-3树
	if !node.Left.Is23() {
		return false
	}
	// 右子树也要判断是否是遵循2-3树
	if !node.Right.Is23() {
		return false
	}
	return true
}

// IsBalanced 节点所在的子树是否平衡，是否有 blackNum 个黑链接
func (node *LLRBTNode) IsBalanced(blackNum int) bool {
	if node == nil {
		return blackNum == 0
	}
	if !IsRed(node) {
		blackNum = blackNum - 1
	}
	if !node.Left.IsBalanced(blackNum) {
		fmt.Println("node.Left to leaf black link is not ", blackNum)
		return false
	}
	if !node.Right.IsBalanced(blackNum) {
		fmt.Println("node.Right to leaf black link is not ", blackNum)
		return false
	}
	return true
}

// 验证ces
func main() {
	tree := NewLLRBTree()
	values := []int64{2, 3, 7, 10, 10, 10, 10, 23, 9, 102, 109, 111, 112, 113}
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

	tree.MidOrder()

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

	if tree.IsLLRBTree() {
		fmt.Println("is a llrb tree")
	} else {
		fmt.Println("is not llrb tree")
	}
}

/*
程序是递归程序，如果改写为非递归形式，效率和性能会更好，
在此就不实现了，理解左倾红黑树添加和删除的总体思路即可

应用场景
红黑树可以用来作为字典 Map 的基础数据结构，可以存储键值对，然后通过一个键，
可以快速找到键对应的值，相比哈希表查找，不需要占用额外的空间。
我们以上的代码实现只有 value，没有 key:value，可以简单改造实现字典。

Java 语言基础类库中的 HashMap，TreeSet，TreeMap 都有使用到，
C++ 语言的 STL 标准模板库中，map 和 set 类也有使用到。
很多中间件也有使用到，比如 Nginx，
但 Golang 语言标准库并没有它

最后，上述应用场景使用的红黑树都是普通红黑树，并不是本文所介绍的左倾红黑树
左倾红黑树作为红黑树的一个变种，只是被设计为更容易理解而已，变种只能是变种，
# 工程上使用得更多的还是普通红黑树，所以我们仍然需要学习普通的红黑树
 */