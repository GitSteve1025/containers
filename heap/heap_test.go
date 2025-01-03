// Copyright (c) 2024 Tecy.
// This file is licensed under the MIT License.
// See the LICENSE file in the project root for more information.

package heap

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func same(t *testing.T, left any, right any) {
	if !reflect.DeepEqual(left, right) {
		t.Fatal(left, "is not equal to", right)
	}
}

func TestHeapBasicFunction(t *testing.T) {
	cmp := func(a int, b int) bool {
		return a < b
	}

	heap := New(cmp)
	for i := 10; i >= 0; i-- {
		heap.Push(i)
	}

	for i := 0; i <= 10; i++ {
		same(t, heap.Empty(), false)
		same(t, 11-i, heap.Size())
		same(t, i, heap.Top())
		same(t, i, heap.Pop())
	}

	same(t, heap.Empty(), true)
	same(t, heap.Top(), 0) // default value
	same(t, heap.Pop(), 0) // default value
}

func TestMakeHeap(t *testing.T) {
	cmp := func(a int, b int) bool {
		return a < b
	}

	const N = 10
	var val []int
	for i := 0; i < N; i++ {
		val = append(val, i)
	}

	heap := NewWithData(cmp, val...)
	t.Log(val)
	for i := 0; i < N; i++ {
		same(t, heap.Empty(), false)
		same(t, N-i, heap.Size())
		same(t, i, heap.Top())
		same(t, i, heap.Pop())
	}

	same(t, heap.Empty(), true)
	same(t, heap.Top(), 0) // default value
	same(t, heap.Pop(), 0) // default value
}

func TestBigData(t *testing.T) {
	cmp := func(a int, b int) bool {
		return a < b
	}

	const N = 10000000
	var expect []int
	heap := New(cmp)
	for i := 0; i < N; i++ {
		val := rand.Intn(1000000000)
		expect = append(expect, val)
		heap.Push(val)
	}

	sort.Slice(expect, func(i, j int) bool {
		return expect[i] < expect[j]
	})
	for i := 0; i < N; i++ {
		same(t, expect[i], heap.Pop())
	}
}

func TestEfficiency(t *testing.T) {
	cmp := func(a int, b int) bool {
		return a < b
	}

	const N = 10000000
	var expect []int
	for i := 0; i < N; i++ {
		expect = append(expect, rand.Intn(1000000000))
	}

	startTime := time.Now()
	heap := NewWithData(cmp, expect...) // O(n)
	t.Log("make heap with", N, "digits costs", time.Since(startTime))

	startTime = time.Now()
	for !heap.Empty() {
		heap.Pop()
	}
	t.Log("Pop", N, "digits costs", time.Since(startTime))
}
