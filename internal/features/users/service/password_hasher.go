package users_service

type PasswordHasher interface {
	Hash(password string) (string, error)
	//Compare(password, hash string) (bool, error)
}
