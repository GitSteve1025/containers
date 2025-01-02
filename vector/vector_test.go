// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package vector

import (
	"testing"
	"time"
)

func TestMake(t *testing.T) {
	vec := make(Vector[int], 10)
	t.Log(vec, vec.Size(), vec.Capacity(), vec.Empty())
	vec.Assign(5, 10)
	t.Log(vec, vec.Size(), vec.Capacity(), vec.Empty())
	vec.Resize(3)
	t.Log(vec, vec.Size(), vec.Capacity(), vec.Empty())
}

func TestFrontBack(t *testing.T) {
	vec := make(Vector[int], 10)
	t.Log(*vec.Front(), *vec.Back())
	vec.Clear()
	if vec.Front() != nil || vec.Back() != nil {
		t.Error("front or back is invalid")
	}
}

func TestResize(t *testing.T) {
	vec := NewWithData(1, 2, 3, 4, 5)
	vec.Resize(10)
	for i := range *vec {
		t.Log((*vec)[i])
	}
	vec.Resize(3)
	for i := range *vec {
		t.Log((*vec)[i])
	}
}

func TestPushPop(t *testing.T) {
	vec := make(Vector[int], 5)
	for i := range vec {
		vec[i] = i
	}

	for i := 0; i < 10; i++ {
		vec.PushBack(i)
		t.Log(vec, vec.Size(), vec.Capacity(), vec.Empty())
	}

	for i := 0; i < 15; i++ {
		vec.PopBack()
		t.Log(vec, vec.Size(), vec.Capacity(), vec.Empty())
	}

	if vec.PopBack() != 0 {
		t.Error("PopBack defalut value is invalid")
	}
}

func TestInsertBack(t *testing.T) {
	vec := New[int]()
	for i := 0; i < 10; i++ {
		vec.Insert(i, i)
		t.Log(*vec, vec.Capacity(), vec.Size())
	}
}

func TestPushPopEfficiency(t *testing.T) {
	var vec Vector[int]
	const N = 100000000

	start := time.Now()
	for i := 0; i < N; i++ {
		vec.PushBack(i)

	}
	t.Log("PushBack", N, "val costs", time.Since(start))
	// t.Log(vec)

	start = time.Now()
	for !vec.Empty() {
		vec.PopBack()
	}
	t.Log("PopBack", N, "val costs", time.Since(start))
	// t.Log(vec)
}

func TestInsert(t *testing.T) {
	var vec Vector[int]

	for i := 0; i < 10; i++ {
		vec.Insert(i, i)
		t.Log(vec, vec.Size())
	}

	for i := 0; i < 10; i++ {
		vec.Insert(10-i, i)
		t.Log(vec, vec.Size())
	}
}

func TestErase(t *testing.T) {
	vec := make(Vector[int], 15)
	for i := 0; i < 15; i++ {
		vec[i] = i
	}

	for i := 0; i < 5; i++ {
		t.Log(vec, vec.Erase(0))
	}

	for i := 0; i < 5; i++ {
		t.Log(vec, vec.Erase(4))
	}

	if vec.Erase(100) != 0 {
		t.Error("Erase defalut value is invalid")
	}

	for i := 0; i < 5; i++ {
		t.Log(vec, vec.Erase(vec.Size()-1))
	}
}

func TestAt(t *testing.T) {
	var vec Vector[string]
	vec.PushBack("a")
	vec.PushBack("b")
	vec.PushBack("c")
	for i := 0; i < 5; i++ {
		val := vec.At(i)
		if val != nil {
			t.Log(val)
		} else {
			t.Log("nil")
		}
	}
}

func TestShrinkToFit(t *testing.T) {
	const N = 10
	const B = 5
	vec := make(Vector[string], N)
	for i := 0; i < B; i++ {
		vec.PopBack()
	}
	t.Log(vec.Size(), vec.Capacity())
	vec.ShrinkToFit()
	if N-B != vec.Capacity() {
		t.Fatal("shrink to fit failed")
	}
}

func TestNew(t *testing.T) {
	vec := New[int]()
	for i := 0; i < 10; i++ {
		vec.PushBack(i)
		t.Log(vec.Size())
	}
	for i := 0; i < 10; i++ {
		vec.Insert(i, 0)
		t.Log(vec.Size())
	}
	for i := 0; i < 10; i++ {
		vec.Erase(i)
		t.Log(vec.Size())
	}
	for i := 0; i < 10; i++ {
		vec.PopBack()
		t.Log(vec.Size())
	}
}

func TestNewWithDataAndClear(t *testing.T) {
	vec := NewWithData(1, 2, 3, 4, 5)
	if vec.Size() != 5 {
		t.Error("new with data is wrong")
	}

	for i := 0; i < 5; i++ {
		if vec.PopBack() != 5-i {
			t.Error("popback data is invalid")
		}
	}

	vec.Clear()
	if !vec.Empty() {
		t.Error("clear is invalid")
	}
	vec.Clear()
	vec.Clear()
	vec.Assign(1, 1)
	vec.Clear()
}
