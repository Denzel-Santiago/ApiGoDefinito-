package application

type PasswordService interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, plainPassword string) bool
}
