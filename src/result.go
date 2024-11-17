package result_or

// Success represents a success result.
type Success struct{}

// Created represents a created result.
type Created struct{}

// Deleted represents a deleted result.
type Deleted struct{}

// Updated represents an updated result.
type Updated struct{}

// Result contains static-like properties for each result type.
var Result = struct {
	Success Success
	Created Created
	Deleted Deleted
	Updated Updated
}{
	Success: Success{},
	Created: Created{},
	Deleted: Deleted{},
	Updated: Updated{},
}
