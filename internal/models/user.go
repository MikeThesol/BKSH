package models

type User struct {
	ID           uint64
	Username     string
	Login        string
	HashPassword string
	AvatarURL    string
	Status       string
}

type UserResponse struct {
	User
	Books []Book
}
