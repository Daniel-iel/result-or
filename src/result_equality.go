package result_or

import (
	"hash/fnv"
	"reflect"
)

func (e Error) Equals(other Error) bool {
	return e.Message == other.Message
}

type ErrorOr[T any] struct {
	value  *T
	errors []Error
}

func (eo ErrorOr[T]) IsError() bool {
	return eo.value == nil
}

func (eo ErrorOr[T]) Equals(other ErrorOr[T]) bool {
	if !eo.IsError() {
		return !other.IsError() && reflect.DeepEqual(eo.value, other.value)
	}

	return other.IsError() && checkIfErrorsAreEqual(eo.errors, other.errors)
}

func (eo ErrorOr[T]) HashCode() uint64 {
	h := fnv.New64a()

	if !eo.IsError() {
		h.Write([]byte(reflect.TypeOf(*eo.value).String()))
		return h.Sum64()
	}

	for _, err := range eo.errors {
		h.Write([]byte(err.Message))
	}

	return h.Sum64()
}

func checkIfErrorsAreEqual(errors1, errors2 []Error) bool {
	if len(errors1) != len(errors2) {
		return false
	}

	for i := range errors1 {
		if !errors1[i].Equals(errors2[i]) {
			return false
		}
	}

	return true
}
