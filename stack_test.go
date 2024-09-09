package stack

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}
}

func TestStack_Pop(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, err := s.Pop()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != 3 {
		t.Errorf("expected 3, got %d", val)
	}
	val, err = s.Pop()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != 2 {
		t.Errorf("expected 2, got %d", val)
	}
	val, err = s.Pop()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != 1 {
		t.Errorf("expected 1, got %d", val)
	}
	_, err = s.Pop()
	if err == nil {
		t.Errorf("expected error on empty stack, got none")
	}
}

func TestStack_Peek(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	val, err := s.Peek()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != 2 {
		t.Errorf("expected 2, got %d", val)
	}
	val, err = s.Peek()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != 2 {
		t.Errorf("expected 2, got %d", val)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := Stack[int]{}

	if !s.IsEmpty() {
		t.Errorf("expected stack to be empty")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("expected stack to not be empty")
	}
	_, _ = s.Pop()
	if !s.IsEmpty() {
		t.Errorf("expected stack to be empty after popping all elements")
	}
}

func TestStack_Size(t *testing.T) {
	s := Stack[int]{}
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}
	_, _ = s.Pop()
	if s.Size() != 1 {
		t.Errorf("expected size 1 after pop, got %d", s.Size())
	}
}

func BenchmarkStack_Push(b *testing.B) {
	stack := Stack[int]{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	stack := Stack[int]{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.ResetTimer() // Reset the timer to benchmark only the Pop operation
	for i := 0; i < b.N; i++ {
		_, _ = stack.Pop()
	}
}

func BenchmarkStack_PushPop(b *testing.B) {
	stack := Stack[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
		_, _ = stack.Pop()
	}
}

func BenchmarkStack_Peek(b *testing.B) {
	stack := Stack[int]{}
	stack.Push(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = stack.Peek()
	}
}

func BenchmarkStack_Size(b *testing.B) {
	stack := Stack[int]{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Size()
	}
}

func BenchmarkStack_IsEmpty(b *testing.B) {
	stack := Stack[int]{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.IsEmpty()
	}
}
