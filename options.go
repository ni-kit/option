package option

type (
	EachFunc[T any]         func(some T)
	EachFunci[T any]        func(i int, some T)
	FilterFunc[T any]       func(some T) bool
	FilterFunci[T any]      func(i int, some T) bool
	MapFunc[T any, R any]   func(some T) R
	MapFunci[T any, R any]  func(i int, some T) R
	FoldFunc[T any, R any]  func(acc R, next T) R
	FoldFunci[T any, R any] func(i int, acc R, next T) R
	// temporary due to the limited nature of current generics implementation
	MaptFunc[T any]  func(some T) T
	MaptFunci[T any] func(i int, some T) T
)

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

func (opts Options[T]) FilterIdx(fn FilterFunci[T]) []T {
	return Filteri(opts, fn)
}

func (opts Options[T]) OFilteri(fn FilterFunci[T]) Options[T] {
	return OFilteri(opts, fn)
}

func (opts Options[T]) FilterPtri(fn FilterFunci[*T]) []T {
	return FilterPtri(opts, fn)
}

func (opts Options[T]) OFilterPtri(fn FilterFunci[*T]) Options[T] {
	return OFilterPtri(opts, fn)
}

func (opts Options[T]) Each(fn EachFunc[T]) {
	Each(opts, fn)
}

func (opts Options[T]) EachPtr(fn EachFunc[*T]) {
	EachPtr(opts, fn)
}

func (opts Options[T]) Eachi(fn EachFunci[T]) {
	Eachi(opts, fn)
}

func (opts Options[T]) EachPtri(fn EachFunci[*T]) {
	EachPtri(opts, fn)
}

func (opts *Options[T]) Push(t T) {
	*opts = append(*opts, O(t))
}

func (opts Options[T]) Append(t T) Options[T] {
	return append(opts, O(t))
}

func (opts Options[T]) Mapt(fn MaptFunc[T]) (res []T) {
	for _, opt := range opts {
		opt.Some(func(some T) {
			res = append(res, fn(some))
		})
	}
	return res
}

func (opts Options[T]) Mapti(fn MaptFunci[T]) (res []T) {
	for i, opt := range opts {
		opt.Some(func(some T) {
			res = append(res, fn(i, some))
		})
	}
	return res
}

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

func Eachi[T any](opts Options[T], fn EachFunci[T]) {
	for i, opt := range opts {
		opt.Some(func(some T) {
			fn(i, some)
		})
	}
}

func EachPtri[T any](opts Options[T], fn EachFunci[*T]) {
	for i, opt := range opts {
		opt.SomePtr(func(some *T) {
			fn(i, some)
		})
	}
}

// Map is used to iterate through Options and create a new slice from Some values
func Map[T any, R any](opts Options[T], fn MapFunc[T, R]) (res []R) {
	return Foldl(opts, func(res []R, next T) []R {
		return append(res, fn(next))
	}, []R{})
}

// Mapi is used to iterate through Options and create a new slice from Some values
// allows accessing index of an element
func Mapi[T any, R any](opts Options[T], fn MapFunci[T, R]) (res []R) {
	return Foldli(opts, func(i int, res []R, next T) []R {
		return append(res, fn(i, next))
	}, []R{})
}

// OFilter is used to iterate through Options and create a new slice of Options based on a provided callback condition
func OFilter[T any](opts Options[T], fn FilterFunc[T]) Options[T] {
	return Foldl(opts, func(res Options[T], next T) Options[T] {
		if fn(next) {
			res.Push(next)
		}
		return res
	}, Options[T]{})
}

// Filter is used to iterate through Options and create a new slice of Some values based on a provided callback condition
func Filter[T any](opts Options[T], fn FilterFunc[T]) []T {
	return Foldl(opts, func(res []T, next T) []T {
		if fn(next) {
			res = append(res, next)
		}
		return res
	}, []T{})
}

// OFilterPtr is used to iterate through Options and create a new slice of Options based on a provided callback condition
func OFilterPtr[T any](opts Options[T], fn FilterFunc[*T]) Options[T] {
	return FoldlPtr(opts, func(res Options[T], next *T) Options[T] {
		if fn(next) {
			res.Push(*next)
		}
		return res
	}, Options[T]{})
}

// FilterPtr is used to iterate through Options and create a new slice of Some values based on a provided callback condition
func FilterPtr[T any](opts Options[T], fn FilterFunc[*T]) (res []T) {
	return FoldlPtr(opts, func(res []T, next *T) []T {
		if fn(next) {
			res = append(res, *next)
		}
		return res
	}, []T{})
}

// Filteri is used to iterate through Options and create a new slice of Some values based on a provided callback condition
func Filteri[T any](opts Options[T], fn FilterFunci[T]) (res []T) {
	return Foldli(opts, func(i int, res []T, next T) []T {
		if fn(i, next) {
			res = append(res, next)
		}
		return res
	}, []T{})
}

// OFilteri is used to iterate through Options and create a new slice of Some values based on a provided callback condition
func OFilteri[T any](opts Options[T], fn FilterFunci[T]) (res Options[T]) {
	return Foldli(opts, func(i int, res Options[T], next T) Options[T] {
		if fn(i, next) {
			res.Push(next)
		}
		return res
	}, Options[T]{})
}

// FilterPtri is used to iterate through Options and create a new slice of Some values based on a provided callback condition
func FilterPtri[T any](opts Options[T], fn FilterFunci[*T]) (res []T) {
	return FoldlPtri(opts, func(i int, res []T, next *T) []T {
		if fn(i, next) {
			res = append(res, *next)
		}
		return res
	}, []T{})
}

// OFilterPtri is used to iterate through Options and create a new slice of Options based on a provided callback condition
func OFilterPtri[T any](opts Options[T], fn FilterFunci[*T]) (res Options[T]) {
	return FoldlPtri(opts, func(i int, res Options[T], next *T) Options[T] {
		if fn(i, next) {
			res.Push(*next)
		}
		return res
	}, Options[T]{})
}

// Foldl is used to iterate over the Options and populate the provided R[esulting] value with the help of a callback
func Foldl[T any, R any](opts Options[T], fn FoldFunc[T, R], start R) R {
	for _, opt := range opts {
		opt.Some(func(some T) {
			start = fn(start, some)
		})
	}
	return start
}

// FoldlPtr is used to iterate over the Options and populate the provided R[esulting] value with the help of a callback
func FoldlPtr[T any, R any](opts Options[T], fn FoldFunc[*T, R], start R) R {
	for _, opt := range opts {
		opt.SomePtr(func(some *T) {
			start = fn(start, some)
		})
	}
	return start
}

// Foldli is used to iterate over the Options and populate the provided R[esulting] value with the help of a callback
func Foldli[T any, R any](opts Options[T], fn FoldFunci[T, R], start R) R {
	for i, opt := range opts {
		opt.Some(func(some T) {
			start = fn(i, start, some)
		})
	}
	return start
}

// FoldlPtri is used to iterate over the Options and populate the provided R[esulting] value with the help of a callback
func FoldlPtri[T any, R any](opts Options[T], fn FoldFunci[*T, R], start R) R {
	for i, opt := range opts {
		opt.SomePtr(func(some *T) {
			start = fn(i, start, some)
		})
	}
	return start
}

// Foldr is used to iterate over the Options and populate the provided R[esulting] value with the help of a callback
// Same as Foldl but goes from right to left
func Foldr[T any, R any](opts Options[T], fn FoldFunc[T, R], start R) R {
	for i := len(opts) - 1; i >= 0; i-- {
		opts[i].Some(func(some T) {
			start = fn(start, some)
		})
	}
	return start
}

// FoldrPtr is used to iterate over the Options and populate the provided R[esulting] value with the help of a callback
func FoldrPtr[T any, R any](opts Options[T], fn FoldFunc[*T, R], start R) R {
	for i := len(opts) - 1; i >= 0; i-- {
		opts[i].SomePtr(func(some *T) {
			start = fn(start, some)
		})
	}
	return start
}

// Foldri is used to iterate over the Options and populate the provided R[esulting] value with the help of a callback
// Same as Foldli but goes from right to left
func Foldri[T any, R any](opts Options[T], fn FoldFunci[T, R], start R) R {
	idx := 0
	for i := len(opts) - 1; i >= 0; i-- {
		opts[i].Some(func(some T) {
			start = fn(idx, start, some)
		})
		idx++
	}
	return start
}
