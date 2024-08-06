package mygo

// P is alias for Ptr
func P[T any](v T) *T {
	return Ptr(v)
}

// V is alias for Value
func V[T any](p *T) T {
	return Value(p)
}

func Ptr[T any](v T) *T {
	return &v
}

func Value[T any](p *T) T {
	if p == nil {
		return *new(T)
	}
	return *p
}
