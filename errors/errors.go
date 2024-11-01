package errors

type ErrorType struct {
	value string
}

type SlugError struct {
	err       string
	slug      string
	errorType ErrorType
}

var (
	ErrorTypeAuthentication = ErrorType{"authentication"}
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
)

func (se *SlugError) Error() string {
	return se.err
}
func (se *SlugError) ErrorType() ErrorType {
	return se.errorType
}

func (se *SlugError) Slug() string {
	return se.slug
}

func NewSlugError(err, slug string) SlugError {
	return SlugError{
		err,
		slug,
		ErrorTypeUnknown,
	}
}

func NewAuthorizationError(err, slug string) SlugError {
	return SlugError{
		err,
		slug,
		ErrorTypeAuthentication,
	}
}

func NewIncorrectInputError(err, slug string) SlugError {
	return SlugError{
		err,
		slug,
		ErrorTypeIncorrectInput,
	}
}
