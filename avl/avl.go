package main

import "fmt"

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

func (a *AVLTree) PreOrderTraverse() {
	preOrderTraverse(a.root)
}

func preOrderTraverse(node *AVLNode) {
	if node == nil {
		return
	}
	fmt.Println(node.key)
	preOrderTraverse(node.left)
	preOrderTraverse(node.right)
}

func (a *AVLTree) Insert(key int, val interface{}) *AVLNode {
	return a.root.insert(key, val)
}

func (a *AVLTree) Delete(key int) {
	a.root.delete(key)
}

func (a *AVLTree) Update(oldKey, newKey int, newVal interface{}) {
	root := a.root.search(oldKey)
	root.key = newKey
	root.val = newVal
	root.rebalance()
}

func (a *AVLTree) Search(key int) *AVLNode {
	return a.root.search(key)
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

func (n *AVLNode) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}
func (n *AVLNode) getBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

func (n *AVLNode) recalcHeight() {
	n.height = max(n.left.Height(), n.right.Height()) + 1
}

func (n *AVLNode) rebalance() *AVLNode {
	root := n
	if nil == root {
		return nil
	}
	root.recalcHeight()
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
			greatestNode := n.left.biggest()
			n.key = greatestNode.key
			n.val = greatestNode.val
			n.left.delete(greatestNode.key)

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

func (n *AVLNode) biggest() *AVLNode {
	root := n
	if root.right != nil {
		return root.right.biggest()
	} else {
		return root
	}
}

func (n *AVLNode) smallest() *AVLNode {
	root := n
	if root.left != nil {
		return root.left.smallest()
	} else {
		return root
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
