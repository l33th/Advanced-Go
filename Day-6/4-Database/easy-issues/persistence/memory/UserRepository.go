package memory

import (
	"errors"
	"github.com/Advanced-Go/Day-6/4-Database/easy-issues/domain"
	"github.com/patrickmn/go-cache"
)

const (
	UsersAllKey = "Users:all"
	UserLastId  = "User:lastId"
)

// UserRepository concrete implementation of in-memory db
type UserRepository struct {
	db *cache.Cache
}

func NewUserRepository() *UserRepository {
	db := cache.New(cache.NoExpiration, cache.NoExpiration)
	db.SetDefault(UserLastId, int64(0))
	db.SetDefault(UsersAllKey, []*domain.User{})
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) All() ([]*domain.User, error) {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		return result.([]*domain.User), nil
	} else {
		return nil, errors.New("Empty list")
	}
}

func (r *UserRepository) GetById(id int64) (*domain.User, error) {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		items := result.([]*domain.User)
		for _, user := range items {
			if user.Id == id {
				return user, nil
			}
		}
		return nil, errors.New("Not Found")
	}
	return nil, errors.New("Not Found")
}

func (r *UserRepository) Create(u *domain.User) error {
	id, _ := r.db.IncrementInt64(UserLastId, int64(1))
	u.Id = id

	result, ok := r.db.Get(UsersAllKey)
	if ok {
		result = append(result.([]*domain.User), u)
		r.db.Set(UsersAllKey, result, cache.NoExpiration)
	}

	return nil
}

func (r *UserRepository) Delete(id int64) error {
	result, ok := r.db.Get(UsersAllKey)
	if ok {
		items := result.([]*domain.User)
		for i, user := range items {
			if user.Id == id {
				items = append(items[:i], items[i+1:]...)
				r.db.Set(UsersAllKey, items, cache.NoExpiration)
				return nil
			}
		}
		return errors.New("Not Found")
	}
	return errors.New("Not Found")
}
