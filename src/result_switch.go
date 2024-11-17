package result_or

// Switch executes the appropriate action based on the state of the ResultOr.
func (r ResultOr[T]) Switch(onResult func(T), onError func([]Error)) {
	if r.IsError() {
		err, _ := r.Errors()
		onError(err)
		return
	}
	onResult(r.Value())
}

// SwitchAsync executes the appropriate action based on the state of the ResultOr asynchronously.
func (r ResultOr[T]) SwitchAsync(onResult func(T) error, onError func([]error) error) error {
	if r.IsError {
		return onError(r.Errors)
	}
	return onResult(r.Result)
}

// SwitchFirst executes the appropriate action based on the state of the ResultOr.
func (r ResultOr[T]) SwitchFirst(onResult func(T), onFirstError func(error)) {
	if r.IsError {
		onFirstError(r.Errors[0])
		return
	}
	onResult(r.Result)
}

// SwitchFirstAsync executes the appropriate action based on the state of the ResultOr asynchronously.
func (r ResultOr[T]) SwitchFirstAsync(onResult func(T) error, onFirstError func(error) error) error {
	if r.IsError {
		return onFirstError(r.Errors[0])
	}
	return onResult(r.Result)
}

// ResultOrExtensions provides additional helper functions for ResultOr.
func Switch[T any](r ResultOr[T], onResult func(T), onError func([]error)) {
	r.Switch(onResult, onError)
}

func SwitchAsync[T any](r ResultOr[T], onResult func(T) error, onError func([]error) error) error {
	return r.SwitchAsync(onResult, onError)
}

func SwitchFirst[T any](r ResultOr[T], onResult func(T), onError func(error)) {
	r.SwitchFirst(onResult, onError)
}

func SwitchFirstAsync[T any](r ResultOr[T], onResult func(T) error, onError func(error) error) error {
	return r.SwitchFirstAsync(onResult, onError)
}
