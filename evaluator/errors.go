package evaluator

type syntaxError struct {
	description string
}

func newSyntaxError(description string) *syntaxError {
	return &syntaxError{description: description}
}

func (e *syntaxError) Error() string {
	return e.description
}

type returnError struct {
	value Object
}

func (e *returnError) Error() string {
	return "Unexpected return statement"
}

type continueError struct{}

func (e *continueError) Error() string {
	return "Unexpected continue statement"
}

func (e *continueError) syntaxError() *syntaxError {
	return newSyntaxError(e.Error())
}

type breakError struct{}

func (e *breakError) Error() string {
	return "Unexpected break statement"
}

func (e *breakError) syntaxError() *syntaxError {
	return newSyntaxError(e.Error())
}
