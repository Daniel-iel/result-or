package result_or

// ToResultOr converts a value into a ResultOr containing that value
func ToResultOr[T any](value T) ResultOr[T] {
	return NewResultOr(value, nil)
}

// ToResultOrFromError converts an error into a ResultOr containing that error
func ToResultOrFromError[T any](err error) ResultOr[T] {
	return NewResultOr[T](zeroValue[T](), []error{err})
}

// ToResultOrFromErrors converts a slice of errors into a ResultOr containing those errors
func ToResultOrFromErrors[T any](errs []error) ResultOr[T] {
	if len(errs) == 0 {
		panic("errors cannot be empty")
	}
	return NewResultOr[T](zeroValue[T](), errs)
}

// ToResultOrFromErrorArray converts an array of errors into a ResultOr containing those errors
func ToResultOrFromErrorArray[T any](errs [][2]error) ResultOr[T] {
	if len(errs) == 0 {
		panic("errors cannot be empty")
	}
	errorList := make([]error, len(errs))
	for i, err := range errs {
		errorList[i] = err[1]
	}
	return NewResultOr[T](zeroValue[T](), errorList)
}

// Helper function to provide a zero value of any type
// func zeroValue[T any]() T {
// 	var zero T
// 	return zero
// }
