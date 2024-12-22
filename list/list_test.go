package list

import "testing"

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

	if lt.head != nil || lt.tail != nil {
		t.Error("head or tail is not nil")
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

	if lt.head != nil || lt.tail != nil {
		t.Error("head or tail is not nil")
	}
}
