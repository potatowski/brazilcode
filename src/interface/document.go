package iface

/*
Document defines an interface for handling document-related operations.
It provides methods for validating, formatting, and generating documents.

Methods:

  - IsValid(doc string) error:
    Validates the given document string and returns an error if it is invalid.

  - Format(doc string) (string, error):
    Formats the given document string into a standardized format and returns
    the formatted string along with any potential error.

  - Generate() (string, error):
    Generates a new document string and returns it along with any potential error.
*/
type Document interface {
	IsValid(doc string) error
	Format(doc string) (string, error)
	Generate() (string, error)
}
