package user_errors

import "errors"

var (
	ErrAuthWrong             = errors.New("o login é inválido, cheque as informações")
	ErrUserNameAlrearyExists = errors.New("o nome de usuário já existe")
	ErrEmailAlrearyExists    = errors.New("o email já existe")
	ErrShortUserName         = errors.New("o Nome é muito curto deve ter pelo menos 4 caracteres")
	ErrInvalidID             = errors.New("o ID é inválido")
	ErrEmptyEmail            = errors.New("o email não pode ser vazio")
	ErrInvalidEmail          = errors.New("o email é inválido")
	ErrShortPassword         = errors.New("a senha deve ter pelo menos 8 caracteres")
	ErrInvalidCreatedAt      = errors.New("a data de criação é inválida")
	ErrInvalidUpdatedAt      = errors.New("a data de atualização é inválida")
)
