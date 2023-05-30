package user

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level int
}
