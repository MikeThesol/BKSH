package repository

type Auth interface {
}

type User interface {
}

type Book interface {
}

type Repository struct {
	Auth
	User
	Book
}

func NewRepository() *Repository {
	return &Repository{}
}
