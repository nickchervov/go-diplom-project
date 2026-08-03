package domain

type DomainError struct {
	Code    int
	Message string
}

func (e DomainError) Error() string {
	return e.Message
}

var (
	ErrEmptyTitle          = &DomainError{Code: 400, Message: "title must be not empty"}
	ErrIncorrectDate       = &DomainError{Code: 400, Message: "incorrect date"}
	ErrIncorrectRepeatRule = &DomainError{Code: 400, Message: "incorrect repeat rule"}
	ErrIncorrectId         = &DomainError{Code: 400, Message: "incorrect id"}
	ErrTaskNotFound        = &DomainError{Code: 404, Message: "task not found"}
	ErrIncorrectPassword   = &DomainError{Code: 400, Message: "incorrect password"}
)
