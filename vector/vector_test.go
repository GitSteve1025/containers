// Copyright (c) 2024 Tecy.  All rights reserved.
// license that can be found in the LICENSE file.

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
	t.Log(vec.Front(), vec.Back())
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
}

func TestPushPopEfficiency(t *testing.T) {
	var vec Vector[int]
	const N = 10000000

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
