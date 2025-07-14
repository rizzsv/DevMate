package domain


type UserRepository interface {
	Create(user *User) error
	Update(user User) error
	Delete(id string) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
	GetAll() ([]User, error)
}
