package domain

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// repositoryに依存していない書き方にする場合に利用
// func NewUser(name, email string) *User {
// 	return &User{Name: name, Email: email}
// }

// type UserRepository interface {
// 	FindByID(id int) (*User, error)
// }
