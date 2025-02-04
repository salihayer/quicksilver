package types

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidVersion = errors.New("invalid version")
	ErrMaxChannels    = errors.New("max channels exceeded")
)

// NewMultiError returns an error aggregate using the given map.
func NewMultiError(errors map[string]error) Errors {
	return Errors{errors}
}

// Error represents aggregated errors, contained in a map.
type Errors struct {
	Errors map[string]error
}

func (e Errors) Error() string {
	return e.details(0)
}

func (e Errors) details(d int) string {
	str := "{"
	d++
	for k, v := range e.Errors {
		str += indent(k, v, d)
	}
	d--
	str += fmt.Sprintf("\n%v}", indentString("  ", d))

	return str
}

func indent(k string, v error, d int) string {
	istr := indentString("  ", d)

	switch err := v.(type) {
	case Errors:
		return fmt.Sprintf("\n%v\"%v\": %v", istr, k, err.details(d))
	default:
		return fmt.Sprintf("\n%v\"%v\": \"%v\"", istr, k, v)
	}
}

func indentString(indent string, d int) string {
	istr := ""
	for i := 0; i < d; i++ {
		istr += indent
	}

	return istr
}
