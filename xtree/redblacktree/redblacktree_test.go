// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package redblacktree

import (
	"fmt"
	"math/rand"
	"testing"
)

func stringify[K any, V any](N *Node[K, V]) string {
	return fmt.Sprintf("Key: %v", N.Key)
}

func output[K any, V any](N *Node[K, V], prefix string, isTail bool, result *string) {
	if N.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(N.right, newPrefix, false, result)
	}
	*result += prefix
	if isTail {
		*result += "└── "
	} else {
		*result += "┌── "
	}
	*result += stringify(N) + "\n"
	if N.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(N.left, newPrefix, true, result)
	}
}

// structure returns the structure of the subtree.
func structure[K any, V any](N *Node[K, V]) string {
	var result = "\n"
	output(N, "", false, &result)
	return result
}

func same(t *testing.T, a any, b any) {
	if a != b {
		t.Fatal("error:", a, "!=", b)
	}
}

// check rbt.
// 1. Nodes are red or black.
// 2. NIL nodes (empty leaf nodes) are black
// 3. Children of red nodes are black.
// 4. Same number of black nodes on each path from root to NIL node
func validater[K any, V any](t *testing.T, rbt *RedBlackTree[K, V]) {
	var value []K
	rootToLeaf := make(map[int]int)
	var dfs func(cur *Node[K, V], count int)
	dfs = func(cur *Node[K, V], count int) {
		if cur == nil {
			count++
			rootToLeaf[count]++
			return
		}

		if isRed(cur) {
			if !isBlack(cur.left) {
				t.Error("cur is red but left is not black")
			}

			if !isBlack(cur.right) {
				t.Error("cur is red but right is not black")
			}
		} else {
			count++
		}

		dfs(cur.left, count)
		value = append(value, cur.Key)
		dfs(cur.right, count)
	}
	dfs(rbt.root, 0)

	if len(rootToLeaf) > 1 {
		t.Error("number of black nodes on each path from root to NIL node is not same")
	}

	for i := 1; i < len(value); i++ {
		if rbt.comparator(value[i], value[i-1]) {
			t.Error("rbt is not a balance tree")
		}
	}
}

func TestRotateLeft(t *testing.T) {
	//     |                       |
	//     N                       S
	//    / \     rotateLeft(N)   / \
	//   L   S    ==========>    N   R
	//      / \                 / \
	//     M   R               L   M

	N := &Node[int, int]{Key: 1, Value: 1}
	L := &Node[int, int]{Key: 2, Value: 2}
	S := &Node[int, int]{Key: 3, Value: 3}
	M := &Node[int, int]{Key: 4, Value: 4}
	R := &Node[int, int]{Key: 5, Value: 5}

	N.left = L
	N.right = S
	L.parent = N
	S.parent = N
	S.left = M
	S.right = R
	M.parent = S
	R.parent = S

	rbt := New[int, int](func(left, right int) bool { return left < right })
	rbt.root = N

	t.Log(structure(N))
	rbt.rotateLeft(N)
	t.Log(structure(S))

	same(t, S.left, N)
	same(t, S.right, R)
	same(t, N.parent, S)
	same(t, R.parent, S)
	same(t, N.left, L)
	same(t, N.right, M)
	same(t, L.parent, N)
	same(t, M.parent, N)
	same(t, rbt.root, S)
}

func TestRotateRight(t *testing.T) {
	//       |                   |
	//       N                   S
	//      / \   r-rotate(N)   / \
	//     S   R  ==========>  L   N
	//    / \                     / \
	//   L   M                   M   R
	N := &Node[int, int]{Key: 1, Value: 1}
	L := &Node[int, int]{Key: 2, Value: 2}
	S := &Node[int, int]{Key: 3, Value: 3}
	M := &Node[int, int]{Key: 4, Value: 4}
	R := &Node[int, int]{Key: 5, Value: 5}

	N.left = S
	N.right = R
	S.parent = N
	R.parent = N
	S.left = L
	S.right = M
	L.parent = S
	M.parent = S

	rbt := New[int, int](func(left, right int) bool { return left < right })
	rbt.root = N

	t.Log(structure(N))
	rbt.rotateRight(N)
	t.Log(structure(S))

	same(t, S.left, L)
	same(t, S.right, N)
	same(t, L.parent, S)
	same(t, N.parent, S)
	same(t, N.left, M)
	same(t, N.right, R)
	same(t, M.parent, N)
	same(t, R.parent, N)
	same(t, rbt.root, S)
}

func TestInsert(t *testing.T) {
	rbt := New[int, int](func(a, b int) bool { return a < b })

	for i := 10; i >= 0; i-- {
		rbt.Insert(i, i)
	}

	validater(t, rbt)

	rbt.root = nil

	for i := 0; i <= 10; i++ {
		rbt.Insert(i, i)
	}

	validater(t, rbt)
}

func TestRandInsert(t *testing.T) {
	N := 10000
	rbt := New[int, int](func(a, b int) bool { return a < b })

	for i := 0; i < 100; i++ {
		rbt.Insert(rand.Intn(N), rand.Intn(N))
	}

	validater(t, rbt)
}

func BenchmarkRandInsert(b *testing.B) {
	MaxVal := 1000000000
	rbt := New[int, int](func(a, b int) bool { return a < b })

	for i := 0; i < b.N; i++ {
		rbt.Insert(rand.Intn(MaxVal), i)
	}
}
