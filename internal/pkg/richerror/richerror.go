package richerror

type Kind int

const (
	KindInvalid Kind = 1 + iota
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	operation    string
	WrappedError error
	Message      string
	Kind         Kind
	meta         map[string]interface{}
}

func New(op string) RichError {
	return RichError{operation: op}
}

func (r RichError) Error() string {
	return r.Message
}

func (r RichError) WithMessage(message string) RichError {
	r.Message = message
	return r
}
func (r RichError) WithKind(kind Kind) RichError {
	r.Kind = kind
	return r
}
func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta
	return r
}
func (r RichError) WithError(err error) RichError {
	r.WrappedError = err
	return r
}
func (r RichError) GetKind() Kind {
	if r.Kind != 0 {
		return r.Kind
	}

	re, ok := r.WrappedError.(RichError)
	if !ok {
		return 0
	}

	return re.GetKind()
}

func (r RichError) GetMessage() string {
	if r.Message != "" {
		return r.Message
	}

	re, ok := r.WrappedError.(RichError)
	if !ok {
		return r.WrappedError.Error()
	}
	return re.GetMessage()
}
