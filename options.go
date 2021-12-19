package option

type Options[T any] []Option[T]

func (opts Options[T]) Filter(fn FilterFunc[T]) (res []T) {
	return Filter(opts, fn)
}

func (opts Options[T]) FilterIdx(fn FilterFuncIdx[T]) (res []T) {
	return FilterIdx(opts, fn)
}

func (opts *Options[T]) Push(t T) {
	*opts = append(*opts, O(t))
}

type (
	FilterFunc[T any]         func(some T) bool
	FilterFuncIdx[T any]      func(i int, some T) bool
	MapFunc[T any, R any]     func(some T) R
	MapFuncIdx[T any, R any]  func(i int, some T) R
	FoldFunc[T any, R any]    func(acc R, next T) R
	FoldFuncIdx[T any, R any] func(i int, acc R, next T) R
)

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
