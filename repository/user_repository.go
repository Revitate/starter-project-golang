package repository

import (
	"encoding/json"
	"starter-project/datasource"
	"starter-project/model"
)

type userRepository struct {
	source datasource.KeyValueDataSource
}

type UserRepository interface {
	Get(id string) (*model.User, error)
	Set(id string, user *model.User) error
}

func NewUserRepository(kv datasource.KeyValueDataSource) UserRepository {
	return &userRepository{kv}
}

func (u *userRepository) Get(id string) (*model.User, error) {
	raw, err := u.source.Get(id)
	if err != nil {
		return nil, err
	}
	user := new(model.User)
	err = json.Unmarshal([]byte(raw), user)
	return user, err
}

func (u *userRepository) Set(id string, user *model.User) error {
	raw, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return u.source.Set(id, string(raw))
}