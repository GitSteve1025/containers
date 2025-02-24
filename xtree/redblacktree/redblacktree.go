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

// setRed sets node to red
func (node *Node[K, V]) setRed() {
	node.color = red
}

// setBlack sets node to black
func (node *Node[K, V]) setBlack() {
	node.color = black
}

// isLeft determinates if that is node's left
func (node *Node[K, V]) isLeft(that *Node[K, V]) bool {
	return node.left != nil && node.left == that
}

// isRight determinates if that is node's right
func (node *Node[K, V]) isRight(that *Node[K, V]) bool {
	return node.right != nil && node.right == that
}

// uncle returns the uncle of node. (parent's brother)
func (node *Node[K, V]) uncle() *Node[K, V] {
	if node.parent == nil || node.parent.parent == nil {
		return nil
	}

	if node.parent.parent.isRight(node.parent) {
		return node.parent.parent.left
	} else {
		return node.parent.parent.right
	}
}

// isRed determines if the node is red.
func isRed[K any, V any](node *Node[K, V]) bool {
	return node != nil && node.color == red
}

// isBlack determines if the node is black.
func isBlack[K any, V any](node *Node[K, V]) bool {
	return node == nil || node.color == black
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
	rbt.insertNode(&Node[K, V]{
		color: red,
		rbt:   rbt,

		Key:   key,
		Value: value,
	})
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

// Merge merges that to this red black tree.
func (rbt *RedBlackTree[K, V]) Merge(that *RedBlackTree[K, V]) {
	// to do
}

// rotateLeft rotates N and returns S.
// N / N.right must not be nil.
//     |                       |
//     N                       S
//    / \     rotateLeft(N)   / \
//   L   S    ==========>    N   R
//      / \                 / \
//     M   R               L   M
func (rbt *RedBlackTree[K, V]) rotateLeft(N *Node[K, V]) {
	S := N.right
	N.right = S.left
	if S.left != nil {
		S.left.parent = N
	}
	S.left = N
	S.parent = N.parent
	if N.parent != nil {
		if N.parent.isLeft(N) {
			N.parent.setLeft(S)
		} else {
			N.parent.setRight(S)
		}
	}
	N.parent = S

	// reset root
	if rbt.isRoot(N) {
		rbt.root = S
	}
}

// rotateRight rotates N and returns S.
// N / N.left must not be nil.
//       |                       |
//       N                       S
//      / \   rotateRight(N)    / \
//     S   R  =============>   L   N
//    / \                     / \
//   L   M                   M   R
func (rbt *RedBlackTree[K, V]) rotateRight(N *Node[K, V]) {
	S := N.left
	N.left = S.right
	if S.right != nil {
		S.right.parent = N
	}
	S.right = N
	S.parent = N.parent
	if N.parent != nil {
		if N.parent.isRight(N) {
			N.parent.setRight(S)
		} else {
			N.parent.setLeft(S)
		}
	}
	N.parent = S

	// reset root
	if rbt.isRoot(N) {
		rbt.root = S
	}
}

// isRoot checks whether node is root.
func (rbt *RedBlackTree[K, V]) isRoot(node *Node[K, V]) bool {
	return rbt.root != nil && rbt.root == node
}

// setLeft sets that to node's left.
func (node *Node[K, V]) setLeft(that *Node[K, V]) {
	node.left = that
	that.parent = node
}

// setRight sets that to node's right.
func (node *Node[K, V]) setRight(that *Node[K, V]) {
	node.right = that
	that.parent = node
}

// insertNode inserts node into redblacktree, node is red.
func (rbt *RedBlackTree[K, V]) insertNode(node *Node[K, V]) {
	// case 1
	// redblacktree is empty, directly insert node.
	if rbt.root == nil {
		rbt.root = node
		return
	}

	cur := rbt.root
	var p *Node[K, V] // parent of cur.
	for cur != nil {
		p = cur
		if rbt.comparator(node.Key, cur.Key) {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}

	if rbt.comparator(node.Key, p.Key) {
		p.setLeft(node)
	} else {
		p.setRight(node)
	}

	rbt.maintainAfterInsert(node)
}

//
func (rbt *RedBlackTree[K, V]) maintainAfterInsert(N *Node[K, V]) {
	// N is root.
	if rbt.isRoot(N) {
		return
	}

	// P is the parent of N.
	P := N.parent

	// case 2, 3
	// The current node's parent is the root node.
	//   {P}(root)         {P}(root)
	//   / \         or    / \
	// <N>                   <N>
	// p.s. <X> X is Red, [X] X is black, {X} unknown

	if rbt.isRoot(P) {
		// case 3
		// The parent node P of the current node N is the root node and is red, just color it black.
		//   <P>(root)         <P>(root)
		//   / \         or    / \
		// <N>                   <N>
		// p.s. <X> X is Red, [X] X is black, {X} unknown

		if isRed(P) {
			P.setBlack()
		}
		return
	}

	// G is the grandparent of N.
	G := P.parent
	// U is the uncle of N.
	U := N.uncle()

	if isRed(P) {
		if isRed(U) {
			// case 4
			// The parent node P and the uncle node U of the current node N are both red.
			// 1. Blacken the P, U nodes and redden the G node.
			// 2. Recursively maintain G nodes.
			//     [G]            [G]          [G]           [G]
			//     /  \          /  \          / \           /  \
			//   <P>  <U>  or  <P>  <U>  or  <U>  <P>  or  <U>  <P>
			//   / \           / \                / \           / \
			// <N>               <N>            <N>               <N>
			// p.s. <X> X is Red, [X] X is black, {X} unknown

			// 1.
			P.setBlack()
			U.setBlack()
			G.setRed()

			// 2.
			rbt.maintainAfterInsert(G)
		} else {
			// case 5
			// The current node N is in the opposite direction of the parent node P.
			// The subtree structure needs to be adjusted to case 6 by a rotation operation.
			//     [G]             [G]
			//     /  \            /  \
			//   <P>  [U]   or   [U]  <P>
			//   / \                  / \
			//     <N>              <N>
			// p.s. <X> X is Red, [X] X is black, {X} unknown

			// LR
			//     [G]                        [G]
			//     /  \     rotateLeft(P)     /  \
			//   <P>  [U]   ============>   <N>  [U]
			//   / \                       /  \
			// [1] <N>                    <P> [3]
			//     / \                   /  \
			//   [2] [3]                [1] [2]

			// RL
			//     [G]                          [G]
			//     /  \      rotateRight(P)     /  \
			//   [U]  <P>    ============>    [U]  <N>
			//        / \                          /  \
			//      <N> [3]                      [1]  <P>
			//      / \                               / \
			//    [1] [2]                           [2] [3]

			if G.isLeft(P) {
				if P.isRight(N) {
					rbt.rotateLeft(P) // LR
					N, P = P, N       // swap N P
				}
			} else {
				if P.isLeft(N) {
					rbt.rotateRight(P) // RL
					N, P = P, N        // swap N P
				}
			}

			// case 6
			// The current node N is in the same direction as the parent node P.
			// 1. If N is a left child node then right-handedly grandfather node G,
			//      otherwise left-handedly grandfather node G.
			// 2. Blacken P, redden G.
			//     [G]             [G]
			//     /  \            /  \
			//   <P>  [U]   or   [U]  <P>
			//   / \                  / \
			//  <N>                     <N>
			// p.s. <X> X is Red, [X] X is black, {X} unknown

			// 1.
			if P.isLeft(N) {
				rbt.rotateRight(G)
			} else {
				rbt.rotateLeft(G)
			}

			// 2.
			P.setBlack()
			G.setRed()
		}
	}
}
