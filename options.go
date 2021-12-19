package option

type Options[T any] []Option[T]

func Slice[T any](ts ...T) (res Options[T]) {
	for _, t := range ts {
		res.Push(t)
	}
	return
}

func (opts Options[T]) OFilter(fn FilterFunc[T]) Options[T] {
	return OFilter(opts, fn)
}

func (opts Options[T]) Filter(fn FilterFunc[T]) []T {
	return Filter(opts, fn)
}

func (opts Options[T]) OFilterPtr(fn FilterFunc[*T]) Options[T] {
	return OFilterPtr(opts, fn)
}

func (opts Options[T]) FilterPtr(fn FilterFunc[*T]) []T {
	return FilterPtr(opts, fn)
}

func (opts Options[T]) FilterIdx(fn FilterFuncIdx[T]) []T {
	return FilterIdx(opts, fn)
}

func (opts Options[T]) OFilterIdx(fn FilterFuncIdx[T]) Options[T] {
	return OFilterIdx(opts, fn)
}

func (opts Options[T]) FilterPtrIdx(fn FilterFuncIdx[*T]) []T {
	return FilterPtrIdx(opts, fn)
}

func (opts Options[T]) OFilterPtrIdx(fn FilterFuncIdx[*T]) Options[T] {
	return OFilterPtrIdx(opts, fn)
}

func (opts Options[T]) Each(fn EachFunc[T]) {
	Each(opts, fn)
}

func (opts Options[T]) EachPtr(fn EachFunc[*T]) {
	EachPtr(opts, fn)
}

func (opts Options[T]) EachIdx(fn EachFuncIdx[T]) {
	EachIdx(opts, fn)
}

func (opts Options[T]) EachPtrIdx(fn EachFuncIdx[*T]) {
	EachPtrIdx(opts, fn)
}

func (opts *Options[T]) Push(t T) {
	*opts = append(*opts, O(t))
}

func (opts Options[T]) Append(t T) Options[T] {
	return append(opts, O(t))
}

type (
	EachFunc[T any]           func(some T)
	EachFuncIdx[T any]        func(i int, some T)
	FilterFunc[T any]         func(some T) bool
	FilterFuncIdx[T any]      func(i int, some T) bool
	MapFunc[T any, R any]     func(some T) R
	MapFuncIdx[T any, R any]  func(i int, some T) R
	FoldFunc[T any, R any]    func(acc R, next T) R
	FoldFuncIdx[T any, R any] func(i int, acc R, next T) R
)

func Each[T any](opts Options[T], fn EachFunc[T]) {
	for _, opt := range opts {
		opt.Some(func(some T) {
			fn(some)
		})
	}
}

func EachPtr[T any](opts Options[T], fn EachFunc[*T]) {
	for _, opt := range opts {
		opt.SomePtr(func(some *T) {
			fn(some)
		})
	}
}

func EachIdx[T any](opts Options[T], fn EachFuncIdx[T]) {
	for i, opt := range opts {
		opt.Some(func(some T) {
			fn(i, some)
		})
	}
}

func EachPtrIdx[T any](opts Options[T], fn EachFuncIdx[*T]) {
	for i, opt := range opts {
		opt.SomePtr(func(some *T) {
			fn(i, some)
		})
	}
}

// Map is used to iterate through an options
func Map[T any, R any](opts Options[T], fn MapFunc[T, R]) (res []R) {
	for _, opt := range opts {
		opt.Some(func(some T) {
			res = append(res, fn(some))
		})
	}
	return res
}

func MapIdx[T any, R any](opts Options[T], fn MapFuncIdx[T, R]) (res []R) {
	for i, opt := range opts {
		opt.Some(func(some T) {
			res = append(res, fn(i, some))
		})
	}
	return res
}

func OFilter[T any](opts Options[T], fn FilterFunc[T]) (res Options[T]) {
	for _, opt := range opts {
		opt.Some(func(some T) {
			if fn(some) {
				res = append(res, opt)
			}
		})
	}
	return res
}

func Filter[T any](opts Options[T], fn FilterFunc[T]) (res []T) {
	for _, opt := range opts {
		opt.Some(func(some T) {
			if fn(some) {
				res = append(res, some)
			}
		})
	}
	return res
}

func OFilterPtr[T any](opts Options[T], fn FilterFunc[*T]) (res Options[T]) {
	for _, opt := range opts {
		opt.SomePtr(func(some *T) {
			if fn(some) {
				res = append(res, opt)
			}
		})
	}
	return res
}

func FilterPtr[T any](opts Options[T], fn FilterFunc[*T]) (res []T) {
	for _, opt := range opts {
		opt.SomePtr(func(some *T) {
			if fn(some) {
				res = append(res, *some)
			}
		})
	}
	return res
}

func FilterIdx[T any](opts Options[T], fn FilterFuncIdx[T]) (res []T) {
	for i, opt := range opts {
		opt.Some(func(some T) {
			if fn(i, some) {
				res = append(res, some)
			}
		})
	}
	return res
}

func OFilterIdx[T any](opts Options[T], fn FilterFuncIdx[T]) (res Options[T]) {
	for i, opt := range opts {
		opt.Some(func(some T) {
			if fn(i, some) {
				res = append(res, opt)
			}
		})
	}
	return res
}

func FilterPtrIdx[T any](opts Options[T], fn FilterFuncIdx[*T]) (res []T) {
	for i, opt := range opts {
		opt.SomePtr(func(some *T) {
			if fn(i, some) {
				res = append(res, *some)
			}
		})
	}
	return res
}

func OFilterPtrIdx[T any](opts Options[T], fn FilterFuncIdx[*T]) (res Options[T]) {
	for i, opt := range opts {
		opt.SomePtr(func(some *T) {
			if fn(i, some) {
				res = append(res, opt)
			}
		})
	}
	return res
}

func Foldl[T any, R any](opts Options[T], fn FoldFunc[T, R], start R) R {
	for _, opt := range opts {
		opt.Some(func(some T) {
			start = fn(start, some)
		})
	}
	return start
}

func FoldIdxl[T any, R any](opts Options[T], fn FoldFuncIdx[T, R], start R) R {
	for i, opt := range opts {
		opt.Some(func(some T) {
			start = fn(i, start, some)
		})
	}
	return start
}

func Foldr[T any, R any](opts Options[T], fn FoldFunc[T, R], start R) R {
	for i := len(opts) - 1; i >= 0; i-- {
		opts[i].Some(func(some T) {
			start = fn(start, some)
		})
	}
	return start
}

func FoldIdxr[T any, R any](opts Options[T], fn FoldFuncIdx[T, R], start R) R {
	idx := 0
	for i := len(opts) - 1; i >= 0; i-- {
		opts[i].Some(func(some T) {
			start = fn(idx, start, some)
		})
		idx++
	}
	return start
}
