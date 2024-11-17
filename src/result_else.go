package result_or

import "context"

// Else executa uma função caso o estado seja erro e retorna o resultado.
// Caso contrário, retorna o valor original.

// ElseErrors executa uma função caso o estado seja erro e retorna os erros modificados.
// Caso contrário, retorna o valor original.
func (r *ResultOr[T]) ElseErrors(onError func([]Error) []Error) ResultOr[T] {
	if !r.IsError() {
		return *r
	}

	errorsResult := onError(r.errors)
	return ResultOr[T]{errors: errorsResult}
}

// ElseError retorna um erro específico caso o estado seja erro.
// Caso contrário, retorna o valor original.
func (r *ResultOr[T]) ElseError(errorValue Error) ResultOr[T] {
	if !r.IsError() {
		return *r
	}

	return ResultOr[T]{errors: []Error{errorValue}}
}

// ElseValue executa uma função caso o estado seja erro e retorna o resultado.
// Caso contrário, retorna o valor original.
func (r *ResultOr[T]) ElseValue(onError func([]Error) T) ResultOr[T] {
	if !r.IsError() {
		return *r
	}

	valueResult := onError(r.errors)
	return ResultOr[T]{value: &valueResult}
}

// ElseStaticValue retorna um valor fixo caso o estado seja erro.
// Caso contrário, retorna o valor original.
func (r *ResultOr[T]) ElseStaticValue(staticValue T) ResultOr[T] {
	if !r.IsError() {
		return *r
	}

	return ResultOr[T]{value: &staticValue}
}

func (r *ResultOr[T]) Else(onError func([]Error) Error) ResultOr[T] {
	if !r.IsError() {
		return *r
	}

	errorResult := onError(r.errors)
	return ResultOr[T]{errors: []Error{errorResult}}
}

// Else executa uma função caso o estado seja erro e retorna o resultado.
// func (r *ResultOr[T]) Else(onError func([]Error) T) ResultOr[T] {
// 	if !r.IsError() {
// 		return *r
// 	}

// 	result := onError(r.errors)
// 	return ResultOr[T]{value: &result}
// }

// ElseError executa uma função para gerar um erro, caso o estado seja erro.
// func (r *ResultOr[T]) ElseError(onError func([]Error) Error) ResultOr[T] {
// 	if !r.IsError() {
// 		return *r
// 	}

// 	err := onError(r.errors)
// 	return ResultOr[T]{errors: []Error{err}}
// }

// AsyncElse executa uma função assíncrona caso o estado seja erro.
func AsyncElse[T any](ctx context.Context, errorOr <-chan ResultOr[T], onError func([]Error) T) <-chan ResultOr[T] {
	out := make(chan ResultOr[T])
	go func() {
		defer close(out)
		select {
		case result := <-errorOr:
			if !result.IsError() {
				out <- result
				return
			}
			val := onError(result.errors)
			out <- ResultOr[T]{value: &val}
		case <-ctx.Done():
			return
		}
	}()
	return out
}

// AsyncElseTask executa uma função assíncrona caso o estado seja erro.
func AsyncElseTask[T any](ctx context.Context, errorOr <-chan ResultOr[T], onError func([]Error) <-chan T) <-chan ResultOr[T] {
	out := make(chan ResultOr[T])
	go func() {
		defer close(out)
		select {
		case result := <-errorOr:
			if !result.IsError() {
				out <- result
				return
			}
			val := <-onError(result.errors)
			out <- ResultOr[T]{value: &val}
		case <-ctx.Done():
			return
		}
	}()
	return out
}

// AsyncElseError executa uma função assíncrona para gerar um erro, caso o estado seja erro.
func AsyncElseError[T any](ctx context.Context, errorOr <-chan ResultOr[T], onError func([]Error) <-chan Error) <-chan ResultOr[T] {
	out := make(chan ResultOr[T])
	go func() {
		defer close(out)
		select {
		case result := <-errorOr:
			if !result.IsError() {
				out <- result
				return
			}
			err := <-onError(result.errors)
			out <- ResultOr[T]{errors: []Error{err}}
		case <-ctx.Done():
			return
		}
	}()
	return out
}
