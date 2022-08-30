package avl

type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	key    int         // 键
	val    interface{} // 值
	height int         // 高度
	left   *AVLNode    // 左节点
	right  *AVLNode    // 右节点
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		&AVLNode{0,
			nil,
			0,
			nil,
			nil},
	}
}

func (a *AVLTree) Insert(key int, val interface{}) *AVLNode {
	return a.root.insert(key, val)
}

func (a *AVLTree) Delete(key int) {}

func (a *AVLTree) Update(oldKey, newKey int, newVal interface{}) {

}

func (a *AVLTree) Select(key int) *AVLNode {

}

func (n *AVLNode) insert(key int, val interface{}) *AVLNode {
	if n == nil {
		return &AVLNode{0, "", 0, nil, nil}
	}
	if key < n.key {
		n.left = n.left.insert(key, val)
	} else if key > n.key {
		n.right = n.right.insert(key, val)
	} else {
		n.val = val
	}
	return n.rebalance() // 平衡
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) recalcHeight() {
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
}
func (n *AVLNode) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.getHeight() - n.right.getHeight()
}

func (n *AVLNode) rebalance() *AVLNode {
	root := n
	if nil == root {
		return nil
	}
	balanceFactor := root.getBalanceFactor()
	// LL
	if balanceFactor > 1 && root.left.getBalanceFactor() > 0 {
		return n.rotateRight()
	}
	// LR
	if balanceFactor > 1 && root.left.getBalanceFactor() <= 0 {
		root.left = root.left.rotateLeft()
		return root.rotateRight()
	}
	// RR
	if balanceFactor < -1 && root.right.getBalanceFactor() <= 0 {
		return root.rotateLeft()
	}
	// RL
	if balanceFactor < -1 && root.right.getBalanceFactor() > 0 {
		root.right = root.right.rotateRight()
		return root.rotateLeft()
	}
	return root
}

func (n *AVLNode) rotateLeft() *AVLNode {
	root := n
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	root.recalcHeight()
	newRoot.recalcHeight()
	return newRoot
}

func (n *AVLNode) rotateRight() *AVLNode {
	root := n
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root
	root.recalcHeight()
	newRoot.recalcHeight()
	return newRoot
}

func (n *AVLNode) search(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	}
	if key > n.key {
		return n.right.search(key)
	}
	return n
}

func (n *AVLNode) delete(key int) {
	// 没有左子树和右子树 删自己
	// 如果只有左子树或者只有右子树
	// 既有左子树和右子树怎么操作
	if n == nil {
		return
	}
	if n.key == key {
		if n.left != nil && n.right != nil {
			// 可以从左子树查 也可以从右子树查最小的 or 最大的

		} else if n.left != nil && n.right == nil {
			n = n.left
		} else if n.left == nil && n.right != nil {
			n = n.right
		} else {
			n = nil
		}

	} else if n.key > key {
		n.left.delete(key)
	} else if n.key < key {
		n.right.delete(key)
	}
	n.rebalance() // 再次平衡
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

