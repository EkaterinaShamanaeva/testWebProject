package model

import (
	"github.com/EkaterinaShamanaeva/testWebProject/app/server"
)

type User struct {
	Id      int
	Name    string
	Surname string
}

// GetAllUsers возвращает массив с данными всех пользователей
func GetAllUsers() (users []User, err error) {
	rows, err := server.Db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}

	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Surname)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, err
}

// NewUser конструктор
func NewUser(name, surname string) *User {
	return &User{Name: name, Surname: surname}
}

// GetUserById получает пользователя по id
func GetUserById(userId string) (u User, err error) {
	err = server.Db.QueryRow("SELECT * FROM users WHERE id = ?", userId).Scan(&u.Id,
		&u.Name, &u.Surname)
	if err != nil {
		return u, err
	}
	return
}

// AddUser добавляет нового пользователя
func (u *User) AddUser() (err error) {
	query := "INSERT INTO users (name, surname) VALUES (?, ?)"
	_, err = server.Db.Exec(query, u.Name, u.Surname)
	return
}
