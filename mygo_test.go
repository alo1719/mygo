package mygo

import (
	"testing"
)

func TestPtr(t *testing.T) {
	cases := []any{1, "a", 3.14, [...]int{1, 2, 3}}
	for _, c := range cases {
		if *Ptr(c) != c {
			t.Errorf("Ptr(%v) = %v, want %v", c, *Ptr(c), c)
		}
	}
}

func TestValueString(t *testing.T) {
	p := P("a")
	if V(p) != *p {
		t.Errorf("Value(%v) = %v, want %v", p, V(p), *p)
	}
	p = nil
	if V(p) != "" {
		t.Errorf("Value(%v) = %v, want %v", p, V(p), "")
	}
}

func TestValueStruct(t *testing.T) {
	type S struct {
		i int
	}
	p := &S{42}
	if V(p) != *p {
		t.Errorf("Value(%v) = %v, want %v", p, V(p), *p)
	}
	p = nil
	if V(p).i != 0 {
		t.Errorf("Value(%v) = %v, want %v", p, V(p), S{})
	}
}

func TestInSlice(t *testing.T) {
	s := []int{1, 2, 3}
	if !InSlice(1, s) {
		t.Errorf("In(1, %v) = false, want true", s)
	}
	if InSlice(4, s) {
		t.Errorf("In(4, %v) = true, want false", s)
	}
}

func TestInMap(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	if !InMap(1, m) {
		t.Errorf("In(1, %v) = false, want true", m)
	}
	if InMap(4, m) {
		t.Errorf("In(4, %v) = true, want false", m)
	}
}

func TestMap(t *testing.T) {
	m := &Map[int, string]{}
	m.Put(1, "a")
	m.Put(2, "b")
	if m.Contains(3) {
		t.Errorf("Contains(3) = true, want false")
	}
	if !m.Contains(1) {
		t.Errorf("Contains(1) = false, want true")
	}
	m.Delete(1)
	if m.Contains(1) {
		t.Errorf("Contains(1) = true, want false")
	}
	if v, _ := m.Get(2); v != "b" {
		t.Errorf("Get(2) = %v, want 'b'", v)
	}
	m.Put(2, "bb")
	if v, _ := m.Get(2); v != "bb" {
		t.Errorf("Get(2) = %v, want 'bb'", v)
	}
	m.Put(3, "c")
	if m.Len() != 2 {
		t.Errorf("Len() = %v, want 2", m.Len())
	}
	t.Log(m.Keys())
}

func TestSet(t *testing.T) {
	s := &Set[int]{}
	s.Add(1)
	s.Add(2)
	if s.Contains(3) {
		t.Errorf("Contains(3) = true, want false")
	}
	if !s.Contains(1) {
		t.Errorf("Contains(1) = false, want true")
	}
	s.Delete(1)
	if s.Contains(1) {
		t.Errorf("Contains(1) = true, want false")
	}
	s.Add(3)
	if s.Len() != 2 {
		t.Errorf("Len() = %v, want 1", s.Len())
	}
	t.Log(s.Elements())
}
