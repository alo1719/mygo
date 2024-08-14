package mygo

// P is alias for Ptr
func P[T any](v T) *T {
	return Ptr(v)
}

// V is alias for Value
func V[T any](p *T) T {
	return Value(p)
}

// In is alias for InSlice
func In[T comparable](e T, s []T) bool {
	return InSlice(e, s)
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

func InSlice[T comparable](e T, s []T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func InMap[K comparable, V any](k K, m map[K]V) bool {
	_, ok := m[k]
	return ok
}
