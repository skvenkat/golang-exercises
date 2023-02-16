package main

import "errors"

type User struct {
	ID   int
	Name string
}

type Storage struct {
	users []User
}

func NewStorage(users []User) Storage {
	return Storage{
		users: users,
	}
}

func (s *Storage) FindUser(id int) (User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}

	return User{}, errors.New("couldn't find the given user from the storage")
}
