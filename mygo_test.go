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
