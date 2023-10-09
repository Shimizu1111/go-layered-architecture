package repository

import (
	"database/sql"

	"github.gom/Shimizu1111/go-layered-architecture/domain"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

type User struct {
	ID    int
	Name  string
	Email string
}

func (r *UserRepositoryImpl) FindByID(id int) (*domain.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
