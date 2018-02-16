package pkgerrors

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/internal/json"
)

type errorStackPairMarshaler struct {
	errKey   string
	stackKey string
}

func NewErrorStackPairMarshaler(errKey, stackKey string) zerolog.ErrorMarshaler {
	return errorStackPairMarshaler{errKey, stackKey}
}

func (m errorStackPairMarshaler) AppendError(buf []byte, err error) []byte {
	return json.AppendError(json.AppendKey(buf, m.errKey), err)
}

func (m errorStackPairMarshaler) AppendErrorWithStack(buf []byte, err error) []byte {
	buf = m.AppendError(buf, err)
	buf = append(buf, ',')
	buf = append(json.AppendKey(buf, m.stackKey), MarshalStack(err)...)
	return buf
}

type errorWithStackMarshaler struct {
	key   string
	inner errorStackPairMarshaler
}

func NewErrorWithStackMarshaler(key, errKey, stackKey string) zerolog.ErrorMarshaler {
	return errorWithStackMarshaler{
		key:   key,
		inner: errorStackPairMarshaler{errKey, stackKey},
	}
}

func (m errorWithStackMarshaler) AppendError(buf []byte, err error) []byte {
	return m.append(buf, err, false)
}

func (m errorWithStackMarshaler) AppendErrorWithStack(buf []byte, err error) []byte {
	return m.append(buf, err, true)
}

func (m errorWithStackMarshaler) append(buf []byte, err error, addStack bool) []byte {
	buf = append(json.AppendKey(buf, m.key), '{')
	if addStack {
		buf = m.inner.AppendErrorWithStack(buf, err)
	} else {
		buf = m.inner.AppendError(buf, err)
	}
	buf = append(buf, '}')
	return buf
}
