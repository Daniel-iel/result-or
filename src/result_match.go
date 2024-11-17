package result_or

// Match executa a função apropriada com base no estado do ResultOr.
func (r *ResultOr[T]) Match[TNextValue any](onValue func(T) TNextValue, onError func([]Error) TNextValue) TNextValue {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.errors != nil {
		return onError(r.errors)
	}

	return onValue(*r.value)
}

// MatchAsync executa a função apropriada de forma assíncrona.
func (r *ResultOr[T]) MatchAsync[TNextValue any](onValue func(T) TNextValue, onError func([]Error) TNextValue) chan TNextValue {
	resultChan := make(chan TNextValue)

	go func() {
		r.mu.RLock()
		defer r.mu.RUnlock()

		var result TNextValue
		if r.errors != nil {
			result = onError(r.errors)
		} else {
			result = onValue(*r.value)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

// MatchFirst executa a função com o primeiro erro, caso exista, ou com o valor.
func (r *ResultOr[T]) MatchFirst[TNextValue any](onValue func(T) TNextValue, onFirstError func(Error) TNextValue) TNextValue {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.errors != nil {
		return onFirstError(r.errors[0])
	}

	return onValue(*r.value)
}

// MatchFirstAsync executa a função com o primeiro erro de forma assíncrona.
func (r *ResultOr[T]) MatchFirstAsync[TNextValue any](onValue func(T) TNextValue, onFirstError func(Error) TNextValue) chan TNextValue {
	resultChan := make(chan TNextValue)

	go func() {
		r.mu.RLock()
		defer r.mu.RUnlock()

		var result TNextValue
		if r.errors != nil {
			result = onFirstError(r.errors[0])
		} else {
			result = onValue(*r.value)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

// Match executa a função apropriada de forma síncrona com base no estado do ResultOr.
func (e *ResultOr[T]) Match[TValue, TNextValue any](r ResultOr[TValue], onValue func(TValue) TNextValue, onError func([]Error) TNextValue) TNextValue {
	return r.Match(onValue, onError)
}

// MatchAsync executa a função apropriada de forma assíncrona.
func (e *ResultOr[T]) MatchAsync[TValue, TNextValue any](r ResultOr[TValue], onValue func(TValue) TNextValue, onError func([]Error) TNextValue) chan TNextValue {
	return r.MatchAsync(onValue, onError)
}

// MatchFirst executa a função apropriada com base no estado do ResultOr.
func (e *ResultOr[T]) MatchFirst[TValue, TNextValue any](r ResultOr[TValue], onValue func(TValue) TNextValue, onError func(Error) TNextValue) TNextValue {
	return r.MatchFirst(onValue, onError)
}

// MatchFirstAsync executa a função com o primeiro erro de forma assíncrona.
func (e *ResultOr[T]) MatchFirstAsync[TValue, TNextValue any](r ResultOr[TValue], onValue func(TValue) TNextValue, onError func(Error) TNextValue) chan TNextValue {
	return r.MatchFirstAsync(onValue, onError)
}