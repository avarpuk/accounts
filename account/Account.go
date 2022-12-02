package account

type Account interface {
	ChangeName(newName string) string
	ChangeAge(newAge int) string

	GetAccount()
}
