package user_interfaces_services

type PassEncryptOutput struct {
	PasswordEncrypted *string
}

type PassEncrypt interface {
	EncryptPassword(passwordDecrypt *string) (*PassEncryptOutput, error)
	ValidatePassword(passEncrypted *string, passwordDecrypted *string) (bool, error)
}
