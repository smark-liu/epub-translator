package parser

type Parser interface {
	// Parse returns the translated file Path.
	Parse() (string, error)
}
