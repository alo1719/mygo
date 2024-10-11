package mygo

// An Ordered Map like Python's dict
// Not thread safe (neither is Go's built-in map)
type Map[K comparable, V any] struct {
	l []K
	m map[K]V
}

// An Ordered Set
// Not thread safe
// Contains/Add is O(1) but Delete is O(n)
type Set[T comparable] struct {
	l []T
	m map[T]struct{}
}

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

func (m *Map[K, V]) Put(k K, v V) {
	if m.m == nil {
		m.m = make(map[K]V)
	}
	if InMap(k, m.m) {
		m.m[k] = v
	} else {
		m.l = append(m.l, k)
		m.m[k] = v
	}
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	v, ok := m.m[k]
	return v, ok
}

func (m *Map[K, V]) Keys() []K {
	return m.l
}

func (m *Map[K, V]) Delete(k K) {
	delete(m.m, k)
	for i, e := range m.l {
		if e == k {
			m.l = append(m.l[:i], m.l[i+1:]...)
			break
		}
	}
}

func (m *Map[K, V]) Contains(k K) bool {
	return InMap(k, m.m)
}

func (m *Map[K, V]) Len() int {
	return len(m.l)
}

func (s *Set[T]) Add(e T) {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	if InMap(e, s.m) {
		return
	} else {
		s.l = append(s.l, e)
		s.m[e] = struct{}{}
	}
}

func (s *Set[T]) Contains(e T) bool {
	return InMap(e, s.m)
}

func (s *Set[T]) Delete(e T) {
	delete(s.m, e)
	for i, v := range s.l {
		if v == e {
			s.l = append(s.l[:i], s.l[i+1:]...)
			break
		}
	}
}

func (s *Set[T]) Len() int {
	return len(s.l)
}

func (s *Set[T]) Elements() []T {
	return s.l
}

func ChanOn[T any](ch chan T) bool {
	if _, ok := <-ch; ok {
		return true
	}
	return false
}
