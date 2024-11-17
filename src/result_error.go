package result_or

// ErrorType represents different types of errors.

// Error struct represents an error with various fields like code, description, type, and metadata.
type Error struct {
	Code        string
	Description string
	Type        ErrorType
	Metadata    map[string]interface{}
}

// NewError creates a new Error of a given type with optional metadata.
func NewError(code, description string, errType ErrorType, metadata map[string]interface{}) Error {
	if metadata == nil {
		metadata = make(map[string]interface{})
	}
	return Error{
		Code:        code,
		Description: description,
		Type:        errType,
		Metadata:    metadata,
	}
}

// Failure creates an Error of type Failure.
func ResultOrFailure(code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, Failure, metadata)
}

// Unexpected creates an Error of type Unexpected.
func ResultOrnexpected(code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, Unexpected, metadata)
}

// Validation creates an Error of type Validation.
func ResultOrValidation(code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, Validation, metadata)
}

// Conflict creates an Error of type Conflict.
func ResultOrConflict(code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, Conflict, metadata)
}

// NotFound creates an Error of type NotFound.
func ResultOrNotFound(code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, NotFound, metadata)
}

// Unauthorized creates an Error of type Unauthorized.
func ResultOrUnauthorized(code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, Unauthorized, metadata)
}

// Forbidden creates an Error of type Forbidden.
func ResultOrForbidden(code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, Forbidden, metadata)
}

// Custom creates an Error with a custom type.
func Custom(typeValue int, code, description string, metadata map[string]interface{}) Error {
	return NewError(code, description, ErrorType(typeValue), metadata)
}

// // Equals compares two Error objects for equality.
// func (e Error) Equals(other Error) bool {
// 	if e.Type != other.Type || e.Code != other.Code || e.Description != other.Description {
// 		return false
// 	}

// 	// Compare Metadata
// 	if len(e.Metadata) != len(other.Metadata) {
// 		return false
// 	}

// 	for key, value := range e.Metadata {
// 		if otherValue, ok := other.Metadata[key]; !ok || !reflect.DeepEqual(value, otherValue) {
// 			return false
// 		}
// 	}

// 	return true
// }

// // HashCode generates a hash code for the Error.
// func (e Error) HashCode() int {
// 	hashCode := 0
// 	hashCode = combineHashCodes(hashCode, e.Code)
// 	hashCode = combineHashCodes(hashCode, e.Description)
// 	hashCode = combineHashCodes(hashCode, e.Type)
// 	hashCode = combineHashCodes(hashCode, int(e.Type))

// 	for key, value := range e.Metadata {
// 		hashCode = combineHashCodes(hashCode, key)
// 		hashCode = combineHashCodes(hashCode, value)
// 	}

// 	return hashCode
// }

// // combineHashCodes is a helper function to combine multiple hash codes.
// func combineHashCodes(hashCode, value interface{}) int {
// 	return hashCode*31 + fmt.Sprintf("%v", value)
// }
