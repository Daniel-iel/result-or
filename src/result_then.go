package result_or

func (r ResultOr[T]) Then[TNextValue any](onValue func(T) ResultOr[TNextValue]) ResultOr[TNextValue] {
	if r.IsError {
		return NewResultOr[TNextValue](zeroValue[TNextValue](), r.Errors)
	}
	return onValue(r.Result)
}

func (r ResultOr[T]) ThenDo(action func(T)) ResultOr[T] {
	if r.IsError {
		return r
	}
	action(r.Result)
	return r
}

func (r ResultOr[T]) ThenAsync[TNextValue any](onValue func(T) ResultOr[TNextValue]) chan ResultOr[TNextValue] {
	resultChan := make(chan ResultOr[TNextValue], 1)
	go func() {
		if r.IsError {
			resultChan <- NewResultOr[TNextValue](zeroValue[TNextValue](), r.Errors)
		} else {
			resultChan <- onValue(r.Result)
		}
	}()
	return resultChan
}

func (r ResultOr[T]) ThenDoAsync(action func(T)) chan ResultOr[T] {
	resultChan := make(chan ResultOr[T], 1)
	go func() {
		if !r.IsError {
			action(r.Result)
		}
		resultChan <- r
	}()
	return resultChan
}

func zeroValue[T any]() T {
	var zero T
	return zero
}
