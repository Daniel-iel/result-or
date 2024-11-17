package result_or

import "context"

func (ro ResultOr[T]) FailIf(onValue func(T) bool, error Error) ResultOr[T] {
	if ro.IsError() {
		return ro
	}

	if onValue(ro.Value()) {
		return ResultOr[T]{errors: []Error{error}}
	}

	return ro
}
func (ro ResultOr[T]) Value() T {
	return *ro.value
}

func (ro ResultOr[T]) FailIfWithErrorBuilder(onValue func(T) bool, errorBuilder func(T) Error) ResultOr[T] {
	if ro.IsError() {
		return ro
	}

	if onValue(ro.Value()) {
		return ResultOr[T]{errors: []Error{errorBuilder(ro.Value())}}
	}

	return ro
}

func (ro ResultOr[T]) FailIfAsync(ctx context.Context, onValue func(context.Context, T) (bool, error), error Error) (ResultOr[T], error) {
	if ro.IsError() {
		return ro, nil
	}

	ok, err := onValue(ctx, ro.Value())
	if err != nil {
		return ro, err
	}

	if ok {
		return ResultOr[T]{errors: []Error{error}}, nil
	}

	return ro, nil
}

func (ro ResultOr[T]) FailIfAsyncWithErrorBuilder(
	ctx context.Context,
	onValue func(context.Context, T) (bool, error),
	errorBuilder func(context.Context, T) (Error, error),
) (ResultOr[T], error) {
	if ro.IsError() {
		return ro, nil
	}

	ok, err := onValue(ctx, ro.Value())
	if err != nil {
		return ro, err
	}

	if ok {
		newError, err := errorBuilder(ctx, ro.Value())
		if err != nil {
			return ro, err
		}
		return ResultOr[T]{errors: []Error{newError}}, nil
	}

	return ro, nil
}
