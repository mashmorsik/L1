package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type Users struct {
	users map[string]string
	mu    sync.RWMutex
}

func NewUsers() *Users {
	return &Users{users: make(map[string]string)}
}

func main() {
	id := uuid.New().String()
	users := NewUsers()
	err := users.AddUser(id, "rickan")
	if err != nil {
		return
	}
	fmt.Println(users.users)

	username := users.GetUsername(id)
	fmt.Printf(username)
}

func (u *Users) AddUser(id, username string) error {
	u.mu.RLock()
	_, ok := u.users[id]
	u.mu.RUnlock()
	if ok {
		fmt.Println("User already exists")
		return nil
	}

	u.mu.Lock()
	u.users[id] = username
	u.mu.Unlock()

	return nil
}

func (u *Users) GetUsername(id string) string {
	u.mu.RLock()
	defer u.mu.RUnlock()

	val, ok := u.users[id]
	if ok {
		return val
	}

	return "User not found."
}
