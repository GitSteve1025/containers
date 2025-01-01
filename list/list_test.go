// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package list

import (
	"testing"
)

func TestNew(t *testing.T) {
	lt := New[string]()
	print(lt)
}

func TestPushPopBack(t *testing.T) {
	lt := New[int]()
	for i := 0; i < 10; i++ {
		lt.PushBack(i)
		t.Log(*lt.Front(), *lt.Back())
	}

	for i := 0; i < 10; i++ {
		t.Log(*lt.Front(), *lt.Back())
		t.Log(lt.PopBack())
	}

	if lt.Front() != nil || lt.Back() != nil {
		t.Error("front or back is not nil")
	}
}

func TestPushPopFront(t *testing.T) {
	lt := New[int]()
	for i := 0; i < 5; i++ {
		lt.PushFront(i)
		t.Log(*lt.Front(), *lt.Back())
	}

	for i := 0; i < 5; i++ {
		t.Log(*lt.Front(), *lt.Back(), lt.PopFront())
	}

	if lt.Front() != nil || lt.Back() != nil {
		t.Error("front or back is not nil")
	}
}

func TestForEach(t *testing.T) {
	lt := New[int]()
	for i := 0; i < 5; i++ {
		lt.PushBack(i)
	}

	a := []int{0, 1, 2, 3, 4}
	at := []int{}
	for cur := lt.Front(); cur != nil; cur = cur.Next() {
		at = append(at, cur.Value)
	}

	for i := 0; i < 5; i++ {
		if a[i] != at[i] {
			t.Error("foreach is wrong")
		}
	}

	b := []int{4, 3, 2, 1, 0}
	bt := []int{}
	for cur := lt.Back(); cur != nil; cur = cur.Prev() {
		bt = append(bt, cur.Value)
	}

	for i := 0; i < 5; i++ {
		if b[i] != bt[i] {
			t.Error("rev foreach is wrong")
		}
	}
}

func TestInsertAndErase(t *testing.T) {
	lt := New[int]()
	mid := lt.PushBack(1)
	lt.InsertBefore(0, mid)
	lt.InsertAfter(2, mid)

	a := []int{0, 1, 2}
	at := []int{}
	for cur := lt.Front(); cur != nil; cur = cur.Next() {
		at = append(at, cur.Value)
	}

	for i := 0; i < 3; i++ {
		if a[i] != at[i] {
			t.Error("insert is wrong")
		}
	}

	lt.Erase(mid.Prev())
	lt.Erase(mid.Next())
	if lt.size != 1 || lt.Front().Value != 1 {
		t.Error("erase is wrong")
	}
	lt.Erase(mid)
	if lt.size != 0 {
		t.Error("erase is wrong")
	}
}

func TestWrongInsert(t *testing.T) {
	a := New[int]()
	ap := a.PushBack(1)
	b := New[int]()
	if b.InsertAfter(0, ap) != nil {
		t.Error("insertAfter error")
	}
	if b.InsertBefore(0, ap) != nil {
		t.Error("insertBefore error")
	}
	if b.PopBack() != 0 {
		t.Error("popback is error")
	}
	if b.PopFront() != 0 {
		t.Error("popfront is error")
	}
	if b.Erase(ap) != 0 {
		t.Error("erase is error")
	}
}

func TestMoveToFront(t *testing.T) {
	list := NewWithData(1, 2, 3)
	back := list.Back()
	list.MoveToFront(back)
	// 3 1 2
	if list.Front().Value != 3 {
		t.Error("MoveToFront is invalid")
	}

	back = list.Back()
	list.MoveToFront(back)
	// 2 3 1
	if list.Front().Value != 2 {
		t.Error("MoveToFront is invalid")
	}

	back = list.Back()
	list.MoveToFront(back)
	// 1 2 3
	if list.Front().Value != 1 {
		t.Error("MoveToFront is invalid")
	}

	front := list.Front()
	list.MoveToFront(front) // nothing to do
	if list.Front().Value != 1 {
		t.Error("MoveToFront is invalid")
	}
}

func TestMoveToBack(t *testing.T) {
	list := NewWithData(1, 2, 3)
	front := list.Front()
	list.MoveToBack(front)
	// 2 3 1
	if list.Back().Value != 1 {
		t.Error("MoveToBack is invalid")
	}

	front = list.Front()
	list.MoveToBack(front)
	// 3 1 2
	if list.Back().Value != 2 {
		t.Error("MoveToBack is invalid")
	}

	front = list.Front()
	list.MoveToBack(front)
	// 1 2 3
	if list.Back().Value != 3 {
		t.Error("MoveToBack is invalid")
	}

	back := list.Back()
	list.MoveToBack(back) // nothing to do
	if list.Back().Value != 3 {
		t.Error("MoveToBack is invalid")
	}
}

func TestMove(t *testing.T) {
	list := NewWithData(1, 2, 3, 4)
	a := list.Front()
	b := list.Front().Next()
	c := list.Front().Next().Next()
	d := list.Front().Next().Next().Next()

	list.MoveBefore(d, c)
	// 1 2 4 3
	if list.Back().Value != 3 {
		t.Error("MoveBefore is invalid")
	}

	list.MoveBefore(c, a)
	// 3 1 2 4
	if list.Front().Value != 3 {
		t.Error("MoveBefore is invalid")
	}

	list.MoveBefore(c, c) // nothing to do
	// 3 1 2 4
	if list.Front().Value != 3 {
		t.Error("MoveBefore is invalid")
	}

	list.MoveAfter(b, d)
	// 3 1 4 2
	if list.Back().Value != 2 {
		t.Error("MoveAfter is invalid")
	}

	list.MoveAfter(a, b)
	// 3 4 2 1
	if list.Back().Value != 1 {
		t.Error("MoveAfter is invalid")
	}

	list.MoveAfter(d, d) // nothing to do
	// 3 4 2 1
	if list.Back().Value != 1 {
		t.Error("MoveAfter is invalid")
	}

	list.move(a, a) // nothing to do
	// 3 4 2 1
	if list.Back().Value != 1 {
		t.Error("MoveAfter is invalid")
	}

	if list.PopBack() != 1 {
		t.Error("Move is invalid")
	}
	if list.PopBack() != 2 {
		t.Error("Move is invalid")
	}
	if list.PopBack() != 4 {
		t.Error("Move is invalid")
	}
	if list.PopBack() != 3 {
		t.Error("Move is invalid")
	}
}

func TestAllFunction(t *testing.T) {
	list := NewWithData(1, 2, 3, 4, 5)
	if list.Front().Value != 1 || list.Back().Value != 5 {
		t.Error("list front or back is invalid")
	}
	if list.Size() != 5 || list.Empty() {
		t.Error("NewWithData is invalid")
	}
	list.Clear()
	if list.Size() != 0 || !list.Empty() {
		t.Error("clear is empty")
	}
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			list.PushFront(i)
		} else {
			list.PushBack(i)
		}
	}
	// list -> 4 2 0 1 3

	for i := 4; i >= 0; i-- {
		if i%2 == 0 {
			if list.Front().Value != i {
				t.Error("list push front is invalid")
			}
			list.PopFront()
		} else {
			if list.Back().Value != i {
				t.Error("list push back is invalid")
			}
			list.PopBack()
		}
	}

	if list.PopFront() != 0 {
		t.Error("popfront default value is invalid")
	}

	if list.PopBack() != 0 {
		t.Error("popback default value is invalid")
	}

	temp := NewWithData(1)
	if list.Erase(temp.Front()) != 0 {
		t.Error("erase default value is invalid")
	}

	list.PushBack(1)
	list.PushBack(2)
	it := list.Back()
	list.PushBack(3)
	if list.Front().Value != 1 || list.Back().Value != 3 || list.Size() != 3 {
		t.Error("push back is invalid")
	}
	if list.Erase(it) != 2 {
		t.Error("erase is invalid")
	}
}
