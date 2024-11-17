package result_or

import (
	"fmt"
)

// NewResultOrFromValue cria um ResultOr de um valor.
func NewResultOrFromValue[T any](value T) ResultOr[T] {
	return ResultOr[T]{value: &value}
}

// NewResultOrFromError cria um ResultOr de um erro.
func NewResultOrFromError[T any](err Error) ResultOr[T] {
	return ResultOr[T]{errors: []Error{err}}
}

// NewResultOrFromErrors cria um ResultOr de uma lista de erros.
func NewResultOrFromErrors[T any](errors []Error) (ResultOr[T], error) {
	if errors == nil {
		return ResultOr[T]{}, fmt.Errorf("errors cannot be nil")
	}
	if len(errors) == 0 {
		return ResultOr[T]{}, fmt.Errorf("errors cannot be empty")
	}

	return ResultOr[T]{errors: errors}, nil
}

// NewResultOrFromErrorArray cria um ResultOr a partir de um array de erros.
func NewResultOrFromErrorArray[T any](errors []Error) (ResultOr[T], error) {
	if errors == nil {
		return ResultOr[T]{}, fmt.Errorf("errors cannot be nil")
	}
	if len(errors) == 0 {
		return ResultOr[T]{}, fmt.Errorf("errors cannot be empty")
	}

	return ResultOr[T]{errors: errors}, nil
}
