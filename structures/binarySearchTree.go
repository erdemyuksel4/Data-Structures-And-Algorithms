package structures

type BinarySearchTree struct {
	value       int
	left, right *BinarySearchTree
}

func (tree *BinarySearchTree) Right() *BinarySearchTree { return tree.right }
func (tree *BinarySearchTree) Left() *BinarySearchTree  { return tree.left }
func (tree *BinarySearchTree) Value() int               { return tree.value }
func (tree *BinarySearchTree) NewBinarySearchTree(value int) {
	tree = &BinarySearchTree{
		value: value,
		left:  nil,
		right: nil,
	}
}
func (tree *BinarySearchTree) SetValue(v int) {
	tree.value = v
}
func (tree *BinarySearchTree) Insert(value int) {
	if tree == nil {
		return
	}
	if value < tree.value {
		if tree.left == nil {
			tree.left = &BinarySearchTree{value: value}
			return
		} else {
			tree.left.Insert(value)
			return
		}
	} else {
		if tree.right == nil {
			tree.right = &BinarySearchTree{value: value}
			return
		} else {
			tree.right.Insert(value)
			return
		}
	}
}

func (tree *BinarySearchTree) Search(num int) *BinarySearchTree {
	if tree == nil {
		return nil
	}
	if num < tree.value {
		return tree.left.Search(num)
	} else if num > tree.value {
		return tree.right.Search(num)
	}
	return tree
}
