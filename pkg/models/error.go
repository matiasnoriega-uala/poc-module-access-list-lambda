package models

type InternalError struct {
}

func (e InternalError) Error() string {
	return "Internal Error"
}
