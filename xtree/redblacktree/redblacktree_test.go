// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package redblacktree

import (
	"fmt"
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

func TestRotateLeft(t *testing.T) {
	//     |                       |
	//     N                       S
	//    / \     rotateLeft(N)   / \
	//   L   S    ==========>    N   R
	//      / \                 / \
	//     M   R               L   M

	N := &Node[int, int]{Key: 1, Value: 1, size: 5}
	L := &Node[int, int]{Key: 2, Value: 2, size: 1}
	S := &Node[int, int]{Key: 3, Value: 3, size: 3}
	M := &Node[int, int]{Key: 4, Value: 4, size: 1}
	R := &Node[int, int]{Key: 5, Value: 5, size: 1}

	N.left = L
	N.right = S
	L.parent = N
	S.parent = N
	S.left = M
	S.right = R
	M.parent = S
	R.parent = S

	t.Log(structure(N))
	t.Log(structure(rotateLeft(N)))

	same(t, S.left, N)
	same(t, S.right, R)
	same(t, N.parent, S)
	same(t, R.parent, S)
	same(t, N.left, L)
	same(t, N.right, M)
	same(t, L.parent, N)
	same(t, M.parent, N)

	same(t, S.size, 5)
	same(t, N.size, 3)
	same(t, R.size, 1)
	same(t, L.size, 1)
	same(t, M.size, 1)
}

func TestRotateRight(t *testing.T) {
	//       |                   |
	//       N                   S
	//      / \   r-rotate(N)   / \
	//     S   R  ==========>  L   N
	//    / \                     / \
	//   L   M                   M   R
	N := &Node[int, int]{Key: 1, Value: 1, size: 5}
	L := &Node[int, int]{Key: 2, Value: 2, size: 1}
	S := &Node[int, int]{Key: 3, Value: 3, size: 3}
	M := &Node[int, int]{Key: 4, Value: 4, size: 1}
	R := &Node[int, int]{Key: 5, Value: 5, size: 1}

	N.left = S
	N.right = R
	S.parent = N
	R.parent = N
	S.left = L
	S.right = M
	L.parent = S
	M.parent = S

	t.Log(structure(N))
	t.Log(structure(rotateRight(N)))

	same(t, S.left, L)
	same(t, S.right, N)
	same(t, L.parent, S)
	same(t, N.parent, S)
	same(t, N.left, M)
	same(t, N.right, R)
	same(t, M.parent, N)
	same(t, R.parent, N)

	same(t, S.size, 5)
	same(t, N.size, 3)
	same(t, R.size, 1)
	same(t, L.size, 1)
	same(t, M.size, 1)
}
