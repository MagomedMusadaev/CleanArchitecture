package repositories

import (
	"CleanArchitecture/internal/models"
	"errors"
)

type InMemoryDB struct {
	users map[int]models.User
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		users: make(map[int]models.User),
	}
}

func (db *InMemoryDB) Save(user models.User) error {
	if _, ok := db.users[user.ID]; ok {
		return errors.New("user already exists")
	}
	db.users[user.ID] = user
	return nil
}

func (db *InMemoryDB) FindByID(id int) (models.User, error) {
	user, ok := db.users[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (db *InMemoryDB) DeleteByID(id int) error {
	_, ok := db.users[id]
	if !ok {
		return errors.New("user not found")
	}
	delete(db.users, id)
	return nil
}
