package splat

import "context"

// User is a struct that represents a user. It contains a user id, a user name, and a user email.
type User struct {
	Id    string
	Name  string
	Email string
}

// UserStore is an interface that represents a user store and contains CRUD methods that interact with the user store.
type UserStore interface {
	GetUser(ctx context.Context, id string) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
}
