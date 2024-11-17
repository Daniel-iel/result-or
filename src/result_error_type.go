package result_or

type ErrorType int

const (
	Failure ErrorType = iota
	Unexpected
	Validation
	Conflict
	NotFound
	Unauthorized
	Forbidden
)
