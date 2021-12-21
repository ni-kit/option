package option

import "errors"

type (
	// Ok funcs
	OkFunc[T any]  func(T)
	OkFuncv[T any] func(T) T
	// Err funcs
	ErrFunc  func(error)
	ErrFuncv func(error) error
)

type Result[T any] struct {
	t *T
	e error
}

// ErrNotOK is a therapeutic default error message
var ErrNotOK = errors.New("result is not ok, but it's ok")

func Ok[T any](v T) Result[T] {
	if &v == nil {
		return Result[T]{e: ErrNotOK}
	}
	return Result[T]{t: &v}
}

func Err[T any](err error) Result[T] {
	if err == nil {
		return Result[T]{e: ErrNotOK}
	}
	return Result[T]{e: err}
}

func (r Result[T]) IsOk() bool {
	return r.t != nil
}

func (r Result[T]) IsErr() bool {
	return r.e != nil
}

func (r Result[T]) Switch(
	ok OkFunc[T],
	err ErrFunc,
) {
	if r.IsOk() {
		ok(*r.t)
	} else if r.IsErr() {
		err(r.e)
	}
}

func (r Result[T]) Switchv(
	ok OkFuncv[T],
	err ErrFuncv,
) (T, error) {
	if r.IsOk() {
		return ok(*r.t), nil
	}
	var v T
	if r.IsErr() {
		return v, err(r.e)
	}
	return v, ErrNotOK
}

func (r Result[T]) SwitchErr(
	ok OkFunc[T],
	err ErrFuncv,
) error {
	if r.IsOk() {
		ok(*r.t)
		return nil
	} else if !r.IsErr() {
		return ErrNotOK
	}
	return err(r.e)
}

func (r Result[T]) Must(msg string) T {
	if !r.IsOk() {
		panic(msg)
	}
	return *r.t
}

func (r Result[T]) MustPtr(msg string) *T {
	if !r.IsOk() {
		panic(msg)
	}
	return r.t
}

func (r Result[T]) Mustv(msg string, ok OkFuncv[T]) T {
	if !r.IsOk() {
		panic(msg)
	}
	return ok(*r.t)
}

func (r Result[T]) MustPtrv(msg string, ok OkFunc[*T]) T {
	if !r.IsOk() {
		panic(msg)
	}
	ok(r.t)
	return *r.t
}

func (r Result[T]) Ok(ok OkFunc[T]) {
	if r.IsOk() {
		ok(*r.t)
	}
}

func (r Result[T]) OkErr(ok OkFunc[T]) error {
	if r.IsOk() {
		ok(*r.t)
	}
	return r.e
}

func (r Result[T]) OkPtr(ok OkFunc[*T]) {
	if r.IsOk() {
		ok(r.t)
	}
}

func (r Result[T]) OkPtrt(ok OkFuncv[*T]) *T {
	if r.IsOk() {
		return ok(r.t)
	}
	return nil
}

func (r Result[T]) Default(def T) T {
	if r.IsOk() {
		return *r.t
	}
	return def
}

func (r Result[T]) Defaultv(def T, ok OkFuncv[T]) T {
	if r.IsOk() {
		return ok(*r.t)
	}
	return def
}

func (r Result[T]) DefaultPtrv(def T, ok OkFunc[*T]) T {
	if r.IsOk() {
		ok(r.t)
		return *r.t
	}
	return def
}
