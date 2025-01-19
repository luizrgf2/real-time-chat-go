package user_interfaces_services

type JWTService interface {
	Encode(payload interface{}) (*string, error)
	Decode(token *string) (*interface{}, error)
}
