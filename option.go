package option

type None[T any] struct{}

type Option[T any] struct {
	some *T
}

func (o Option[T]) IsNone() bool {
	return o.some == nil
}

func (o Option[T]) IsSome() bool {
	return !o.IsNone()
}

// Some will execute a function if option is not none
func (o Option[T]) Some(t SomeFunc[T]) bool {
	if o.IsSome() {
		t(*o.some)
		return true
	}
	return false
}

func (o Option[T]) SomePtr(t SomeFunc[*T]) bool {
	if o.IsSome() {
		t(o.some)
		return true
	}
	return false
}

// Switch is used to work with the value in the Option container, returns true if the value is Some
func (o Option[T]) Switch(
	t SomeFunc[T],
	n NoneFunc,
) bool {
	if o.IsSome() {
		t(*o.some)
		return true
	}
	n()
	return false
}

func (o Option[T]) Switcht(
	t SomeFunc[T],
	n NoneFunc,
) *T {
	if o.IsSome() {
		t(*o.some)
		return o.some
	}
	n()
	return nil
}

// Switchv is used to work with the value in the Option container and return it afterwards
func (o Option[T]) Switchv(
	t SomeFuncv[T],
	n NoneFuncv[T],
) T {
	if o.IsSome() {
		return t(*o.some)
	}
	return n()
}

// Default can be used to unpack Option and return either Value or provided default value
func (o Option[T]) Default(
	def T,
) T {
	return o.Switchv(
		func(some T) T {
			return some
		},
		func() T {
			return def
		})
}

// Defaultv can be used to unpack Option and return either Value processed with a callback or provided default value
func (o Option[T]) Defaultv(
	def T,
	t SomeFuncv[T],
) T {
	return o.Switchv(t,
		func() T {
			return def
		})
}

// O is used to construct the Option value
func O[T any](v ...T) Option[T] {
	var t *T
	if len(v) > 0 {
		t = &v[0]
	}
	return Option[T]{
		some: t,
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
