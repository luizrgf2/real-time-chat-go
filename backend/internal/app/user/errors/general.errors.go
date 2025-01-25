package user_errors

import "errors"

var (
	ErrFieldsEmpties = errors.New("todos os campos estão vazios")
)
