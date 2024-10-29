package repository

import (
	"fmt"
	"sync"
)

type User struct {
    ID    string
    Name  string
    Email string
}

type UserRepository interface {
    GetByID(id string) (*User, error)
    Create(name, email string) (*User, error)
    List(page, limit int32) ([]*User, int32, error)
}

type userRepository struct {
    mu    sync.RWMutex
    users map[string]*User
    seq   int
}

func NewUserRepository() UserRepository {
    return &userRepository{
        users: map[string]*User{
            "user1": {ID: "user1", Name: "John Doe", Email: "john@example.com"},
            "user2": {ID: "user2", Name: "Jane Doe", Email: "jane@example.com"},
        },
        seq: 2,
    }
}

func (r *userRepository) GetByID(id string) (*User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    if user, exists := r.users[id]; exists {
        return user, nil
    }
    return nil, fmt.Errorf("user not found")
}

func (r *userRepository) Create(name, email string) (*User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    r.seq++
    id := fmt.Sprintf("user%d", r.seq)
    user := &User{
        ID:    id,
        Name:  name,
        Email: email,
    }
    r.users[id] = user
    return user, nil
}

func (r *userRepository) List(page, limit int32) ([]*User, int32, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    var users []*User
    for _, user := range r.users {
        users = append(users, user)
    }

    start := (page - 1) * limit
    end := start + limit
    total := int32(len(users))

    if start >= int32(len(users)) {
        return []*User{}, total, nil
    }
    if end > int32(len(users)) {
        end = int32(len(users))
    }

    return users[start:end], total, nil
}
