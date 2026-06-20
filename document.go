package brazilcode

import "github.com/potatowski/brazilcode/v3/internal/digit"

// Document defines the interface for Brazilian document operations.
type Document interface {
	// IsValid checks whether the given document string is valid.
	IsValid(doc string) error

	// Format returns the document with standard formatting applied.
	Format(doc string) (string, error)

	// Generate creates a random valid document string.
	// Options can be passed for types that support them (e.g., WithUF for VoterRegistration).
	Generate(opts ...Option) (string, error)
}

// Option configures document generation.
// This is a type alias for the internal option type, ensuring sub-packages
// and root package share the same Option type without circular imports.
type Option = digit.Option

// WithUF sets the state (UF) code for VoterRegistration generation.
//
// Example:
//
//	doc, _ := brazilcode.Generate("VoterRegistration", brazilcode.WithUF("MG"))
func WithUF(uf string) Option {
	return digit.WithUF(uf)
}
