package option

type None[T any] struct{}

type Option[T any] struct {
	some *T
	none None[T]
}

// O is used to construct the Option value
func O[T any](v ...T) Option[T] {
	var t *T
	if len(v) > 0 {
		t = &v[0]
	}
	return Option[T]{
		some: t,
		none: None[T]{},
	}
}

type (
	// Some funcs
	SomeFunc[T any]  func(t T)
	SomeFuncv[T any] func(t T) T
	// None funcs
	NoneFunc         func()
	NoneFuncv[T any] func() T
)

// Switch is used to work with the value in the Option container, returns true if the value is Some
func Switch[T any](
	o Option[T],
	t SomeFunc[T],
	n NoneFunc,
) bool {
	if o.some != nil {
		t(*o.some)
		return true
	}
	n()
	return false
}

// Switcht is used to work with the value in the Option container and return either *some or nil
func Switcht[T any](
	o Option[T],
	t SomeFunc[T],
	n NoneFunc,
) *T {
	if o.some != nil {
		t(*o.some)
		return o.some
	}
	n()
	return o.some
}

// Switchv is used to work with the value in the Option container and return it afterwards
func Switchv[T any](
	o Option[T],
	t SomeFuncv[T],
	n NoneFuncv[T],
) T {
	if o.some != nil {
		return t(*o.some)
	}
	return n()
}

// Default can be used to unpack Option and return either Value or provided default value
func Default[T any](
	o Option[T],
	def T,
) T {
	return Switchv(o,
		func(some T) T {
			return some
		},
		func() T {
			return def
		})
}

// Defaultv can be used to unpack Option and return either Value processed with a callback or provided default value
func Defaultv[T any](
	o Option[T],
	def T,
	t SomeFuncv[T],
) T {
	return Switchv(o, t,
		func() T {
			return def
		})
}
