package errors

type ErrorType struct {
	value string
}

type SlugError struct {
	err        string
	slug       string
	statusCode int
	errorType  ErrorType
}

var (
	ErrorTypeAuthentication = ErrorType{"authentication"}
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
	ErrorTypeNotFound       = ErrorType{"not-found"}
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

func (se *SlugError) StatusCode() int {
	return se.statusCode
}

func NewNotFoundError(err, slug string) *SlugError {
	return &SlugError{
		err,
		slug,
		404,
		ErrorTypeNotFound,
	}
}

func NewSlugError(err, slug string, statusCode int) *SlugError {
	return &SlugError{
		err,
		slug,
		statusCode,
		ErrorTypeUnknown,
	}
}

func NewAuthorizationError(err, slug string, statusCode int) *SlugError {
	return &SlugError{
		err,
		slug,
		statusCode,
		ErrorTypeAuthentication,
	}
}

func NewIncorrectInputError(err, slug string) SlugError {
	return SlugError{
		err,
		slug,
		400,
		ErrorTypeIncorrectInput,
	}
}
