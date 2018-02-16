package zerolog

import (
	"github.com/rs/zerolog/internal/json"
)

// ErrorMarshaler is an interface to customize error message logging.
type ErrorMarshaler interface {
	AppendError(buf []byte, err error) []byte
	AppendErrorWithStack(buf []byte, err error) []byte
}

// NewTextErrorMarshaler() provides a default ErrorMarshaler that
// produces `"error": "an error message"` fragments.
func NewTextErrorMarshaler(name string) ErrorMarshaler {
	return textErrorMarshaler{name}
}

type textErrorMarshaler struct {
	key string
}

func (m textErrorMarshaler) AppendError(buf []byte, err error) []byte {
	return json.AppendError(json.AppendKey(buf, m.key), err)
}

func (m textErrorMarshaler) AppendErrorWithStack(buf []byte, err error) []byte {
	return m.AppendError(buf, err)
}
