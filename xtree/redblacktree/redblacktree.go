// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package redblacktree

// color red = false, black = true
type color bool

const (
	red   color = false
	black color = true
)

type Node[K any, V any] struct {
	size   int // record size of subtree.
	color  color
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
	rbt    *RedBlackTree[K, V] // The red black tree to which this node belongs.

	Key   K
	Value V
}

func (node *Node[K, V]) Prev() *Node[K, V] {
	// to do
	return nil
}

func (node *Node[K, V]) Next() *Node[K, V] {
	// to do
	return nil
}

// Size returns size of subtree.
func (node *Node[K, V]) Size() int {
	return node.size
}

type RedBlackTree[K any, V any] struct {
	comparator func(left K, right K) bool
	root       *Node[K, V]
	size       int
}

func New[K any, V any](comparator func(left K, right K) bool) *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{
		comparator: comparator,
	}
}

func (rbt *RedBlackTree[K, V]) Size() int {
	return rbt.size
}

func (rbt *RedBlackTree[K, V]) Empty() bool {
	return rbt.size == 0
}

func (rbt *RedBlackTree[K, V]) Insert(key K, value V) {
	// to do
}

func (rbt *RedBlackTree[K, V]) Erase(node *Node[K, V]) {
	// to do
}

func (rbt *RedBlackTree[K, V]) LowerBound(key K) *Node[K, V] {
	// to do
	return nil
}

func (rbt *RedBlackTree[K, V]) UpperBound(key K) *Node[K, V] {
	// to do
	return nil
}

func (rbt *RedBlackTree[K, V]) Clear() {
	// to do
}

// Rank returns the rank of the node.
// If the node does not belong to the red black tree, it will return -1.
// node must not be nil.
func (rbt *RedBlackTree[K, V]) Rank(node *Node[K, V]) int {
	if node.rbt != rbt {
		return -1
	}

	// to do

	return 0
}

// Kth returns k-th Node.
// k should be in range [0, size), otherwise Kth returns nil.
func (rbt *RedBlackTree[K, V]) Kth(k int) *Node[K, V] {
	if k < 0 || k >= rbt.size {
		return nil
	}

	// to do

	return nil
}

// Merge merges that to this red black tree.
func (rbt *RedBlackTree[K, V]) Merge(that *RedBlackTree[K, V]) {
	// to do
}

// maintain maintains Node, like size.
// maintain requires bottom-up operation.
// N must not be nil.
func maintain[K any, V any](N *Node[K, V]) {
	N.size = 1
	if N.left != nil {
		N.size += N.left.size
	}

	if N.right != nil {
		N.size += N.right.size
	}
}

// rotateLeft rotates N and returns S.
// N / N.right must not be nil.
//     |                       |
//     N                       S
//    / \     rotateLeft(N)   / \
//   L   S    ==========>    N   R
//      / \                 / \
//     M   R               L   M
func rotateLeft[K any, V any](N *Node[K, V]) *Node[K, V] {
	S := N.right
	N.right = S.left
	if S.left != nil {
		S.left.parent = N
	}
	S.left = N
	S.parent = N.parent
	N.parent = S

	maintain(N)
	maintain(S)
	return S
}

// rotateLeft rotates N and returns S.
// N / N.left must not be nil.
//       |                   |
//       N                   S
//      / \   r-rotate(N)   / \
//     S   R  ==========>  L   N
//    / \                     / \
//   L   M                   M   R
func rotateRight[K any, V any](N *Node[K, V]) *Node[K, V] {
	S := N.left
	N.left = S.right
	if S.right != nil {
		S.right.parent = N
	}
	S.right = N
	S.parent = N.parent
	N.parent = S

	maintain(N)
	maintain(S)
	return S
}
