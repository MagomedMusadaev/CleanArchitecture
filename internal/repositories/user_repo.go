package repositories

import "CleanArchitecture/internal/models"

type UserRepo interface {
	CreateUser(user models.User) error
	GetUserByID(id int) (models.User, error)
	DeleteUserByID(id int) error
}

type InMemoryUserRepository struct {
	db *InMemoryDB
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		db: NewInMemoryDB(),
	}
}

func (repo *InMemoryUserRepository) CreateUser(user models.User) error {
	return repo.db.Save(user)
}

func (repo *InMemoryUserRepository) GetUserByID(id int) (models.User, error) {
	return repo.db.FindByID(id)
}

func (repo *InMemoryUserRepository) DeleteUserByID(id int) error {
	return repo.db.DeleteByID(id)
}
