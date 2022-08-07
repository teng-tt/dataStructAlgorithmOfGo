package main

/*
数与左倾红黑树
红黑树是一种近似平衡的二叉查找树，从 2-3 树或 2-3-4 树衍生而来
通过对二叉树节点进行染色，染色为红或黑节点，来模仿 2-3 树或 2-3-4 树的3节点和4节点
从而让树的高度减小
2-3-4 树对照实现的红黑树是普通的红黑树
2-3   树对照实现的红黑树是一种变种，称为左倾红黑树，其更容易实现

使用平衡树数据结构，可以提高查找元素的速度，
我们用二叉树形式来实现 2-3 树也就是左倾红黑树

2-3 树是一棵严格自平衡的多路查找树
它不是一棵二叉树，是一棵三叉树。具有以下特征：
1、内部节点要么有1个数据元素和2个孩子
2、要么有2个数据元素和3个孩子
3、叶子节点没有孩子，但有1或2个数据元素。
4、所有叶子节点到根节点的长度一致。这个特征保证了完全平衡，非常完美的平衡。
5、每个节点的数据元素保持从小到大排序，两个数据元素之间的子树的所有值大小介于两个数据元素之间。

因为 2-3 树的第二个特征，它是一棵完美平衡的树，非常完美，除了叶子节点，其他的节点都没有空儿子，
所以树的高度非常的小

如果一个内部节点拥有一个数据元素、两个子节点，则此节点为2节点
如果一个内部节点拥有两个数据元素、三个子节点，则此节点为3节点

所有平衡树的核心都在于插入和删除逻辑，主要分析这两个操作


2-3 树插入元素
在插入元素时，需要先找到插入的位置，使用二分查找从上自下查找树节点。
找到插入位置时，将元素插入该位置，然后进行调整，使得满足 2-3 树的特征。主要有三种情况：

插入元素到一个2节点，直接插入即可，这样节点变成3节点。
插入元素到一个3节点，该3节点的父亲是一个2节点，先将节点变成临时的4节点，然后向上分裂调整一次。
插入元素到一个3节点，该3节点的父亲是一个3节点，先将节点变成临时的4节点，然后向上分裂调整，此时
父亲节点变为临时4节点，继续向上分裂调整

核心在于插入3节点后，该节点变为临时4节点，然后进行分裂恢复树的特征
最坏情况为插入节点后，每一次分裂后都导致上一层变为临时4节点，
直到树根节点，这样需要不断向上分裂。

#临时4节点的分裂，细分有六种情况
1、4节点刚好是根节点
2、4节点为左子树，父节点是2节点
3、4节点为左子树，父节点是3节点
4、4节点谓语中间，父节点为3节点
5、4节点为右子树，父节点是2节点
6、4节点为右子树，父节点是3节点
与其他二叉查找树由上而下生长不同，2-3 树是从下至上的生长。

2-3 树删除元素
删除操作就复杂得多了

2-3 树的特征注定它是一棵非常完美平衡的三叉树，其所有子树也都是完美平衡，
所以 2-3 树的某节点的儿子，要么都是空儿子，要么都不是空儿子。
比如 2-3 树的某个节点 A 有两个儿子 B 和 C，儿子 B 和 C 要么都没有孩子，要么孩子都是满的，
不然 2-3 树所有叶子节点到根节点的长度一致这个特征就被破坏了。

基于上面的现实分析删除的不同情况，删除中间节点和叶子节点
情况1：删除中间节点
删除的是非叶子节点，该节点一定是有两棵或者三棵子树的，那么从子树中找到其最小后继节点，该节点是叶子节点，
用该节点替换被删除的非叶子节点，然后再删除这个叶子节点，进入情况2。
如何找到最小后继节点，当有两棵子树时，那么从右子树一直往左下方找，
如果有三棵子树，被删除节点在左边，那么从中子树一直往左下方找，否则从右子树一直往左下方找。

情况2：删除叶子节点
删除的是叶子节点，这时如果叶子节点是3节点，那么直接变为2节点即可，不影响平衡。
如果叶子节点是2节点，那么删除后，其父节点将会缺失一个儿子，破坏了满孩子的 2-3 树特征，
需要进行调整后才能删除。

针对情况2，删除一个2节点的叶子节点，会导致父节点缺失一个儿子，破坏了 2-3 树的特征，可以进行调整变换，
主要有两种调整：
重新分布：尝试从兄弟节点那里借值，然后重新调整节点。
如果被删除的叶子节点有兄弟是3节点，那么从兄弟那里借一个值填补被删除的叶子节点，
然后兄弟和父亲重新分布调整位置

合并：如果兄弟借不到值，合并节点（与父亲的元素），再向上递归处理。
如果兄弟们都是2节点呢，那么就合并节点：将父亲和兄弟节点合并，
如果父亲是2节点，那么父亲就留空了，否则父亲就从3节点变成2节点

如果合并后，父亲节点变空了，也就是说有中间节点留空要怎么办，那么可以继续递归处理
中间节点是空的，那么可以继续从兄弟那里借节点或者和父亲合并，直到根节点，
递归到了根节点后，如果存在空的根节点，我们可以直接把该空节点删除即可，这时树的高度减少一层
 */
