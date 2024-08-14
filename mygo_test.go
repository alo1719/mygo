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
