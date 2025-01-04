package stack

import (
	"math/rand"
	"reflect"
	"testing"
)

func generateData(n int) []int {
	var expect []int
	for i := 0; i < n; i++ {
		expect = append(expect, rand.Intn(1000000000))
	}
	return expect
}

func same(t *testing.T, x, y any) {
	if !reflect.DeepEqual(x, y) {
		t.Fatal(x, "!=", y)
	}
}

func TestStack(t *testing.T) {
	const N = 5
	val := generateData(N)
	stack := New[int]()
	same(t, stack.Empty(), true)
	same(t, stack.Pop(), 0)
	for i := 0; i < N; i++ {
		same(t, stack.Size(), i)
		stack.Push(val[i])
		same(t, *stack.Top(), val[i])
	}
	same(t, stack.Empty(), false)
	for i := 0; i < N; i++ {
		same(t, stack.Size(), N-i)
		same(t, stack.Pop(), val[N-i-1])
	}
	p := stack.Top()
	if p != nil {
		t.Fatal("Top is error")
	}
}

func BenchmarkStack(b *testing.B) {
	val := generateData(b.N)
	b.ResetTimer()
	stack := New[int]()
	for i := 0; i < b.N; i++ {
		stack.Push(val[i])
	}
	for !stack.Empty() {
		stack.Pop()
	}
}
