package result_or

import (
	"errors"
)

type ResultOr[T any] struct {
	value  *T
	errors []Error
}

func NewResultOr[T any](result T, errors []error) ResultOr[T] {
	isError := len(errors) > 0
	return ResultOr[T]{
		Result:  result,
		Errors:  errors,
		IsError: isError,
	}
}

func NewResultOr[T any]() (*ResultOr[T], error) {
	return nil, errors.New("default construction of ResultOr[T] is invalid. Use provided factory methods to instantiate")
}

func FromError[T any](err Error) *ResultOr[T] {
	return &ResultOr[T]{errors: []Error{err}}
}

func FromErrors[T any](errs []Error) (*ResultOr[T], error) {
	if len(errs) == 0 {
		return nil, errors.New("cannot create a ResultOr[T] from an empty collection of errors. Provide at least one error")
	}
	return &ResultOr[T]{errors: errs}, nil
}

func FromValue[T any](value T) *ResultOr[T] {
	return &ResultOr[T]{value: &value}
}

func (r *ResultOr[T]) IsError() bool {
	return len(r.errors) > 0
}

func (r *ResultOr[T]) Errors() ([]Error, error) {
	if !r.IsError() {
		return nil, errors.New("the Errors property cannot be accessed when no errors have been recorded. Check IsError before accessing Errors")
	}
	return r.errors, nil
}

func (r *ResultOr[T]) ErrorsOrEmptyList() []Error {
	if r.IsError() {
		return r.errors
	}
	return []Error{}
}

func (r *ResultOr[T]) GetValue() (T, error) {
	if r.IsError() {
		return *new(T), errors.New("the Value property cannot be accessed when errors have been recorded. Check IsError before accessing Value")
	}
	return *r.value, nil
}

func (r *ResultOr[T]) FirstError() (Error, error) {
	if !r.IsError() {
		return Error{}, errors.New("the FirstError property cannot be accessed when no errors have been recorded. Check IsError before accessing FirstError")
	}
	return r.errors[0], nil
}
