package user_errors

import "errors"

var (
	ErrInvalidID        = errors.New("o ID é inválido")
	ErrEmptyEmail       = errors.New("o email não pode ser vazio")
	ErrInvalidEmail     = errors.New("o email é inválido")
	ErrShortPassword    = errors.New("a senha deve ter pelo menos 8 caracteres")
	ErrInvalidCreatedAt = errors.New("a data de criação é inválida")
	ErrInvalidUpdatedAt = errors.New("a data de atualização é inválida")
)
