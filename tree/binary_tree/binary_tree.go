package binary_tree

import "errors"

type TreeNode struct {
	Value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(input int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.Value == input {
		return errors.New("This node already exist")
	}

	if t.Value > input {
		if t.left == nil {
			t.left = &TreeNode{Value: input}
			return nil
		}

		return t.left.Insert(input)
	}

	if t.Value < input {
		if t.right == nil {
			t.right = &TreeNode{Value: input}
			return nil
		}

		return t.right.Insert(input)
	}

	return nil
}

func (t TreeNode) FindMin() int{
	if t.left == nil {
		return t.Value
	}
	return t.left.FindMin()
}

func (t TreeNode) FindMax() int{
	if t.right == nil {
		return t.Value
	}
	return t.left.FindMin()
}
