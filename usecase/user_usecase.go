package usecase

import (
	"github.gom/Shimizu1111/go-layered-architecture/domain"
	"github.gom/Shimizu1111/go-layered-architecture/repository"
)

// repositoryに依存していない書き方
// type UserUsecase struct {
// 	Repo domain.UserRepository
// }

// repositoryに依存している書き方
type UserUsecase struct {
	Repo *repository.UserRepositoryImpl
}

func (s *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	return s.Repo.FindByID(id)
}
