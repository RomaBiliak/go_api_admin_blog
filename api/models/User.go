package models

import (
	"html"
	"strings"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int64	`json:"id"`
	Name        string	`json:"name"`
	Login		string	`json:"login"`
	Password	string	`json:"password"`
}

var user_id int64 = 0

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Login = html.EscapeString(strings.TrimSpace(u.Login))
	u.Password = strings.Trim(u.Password, " ")
}



func (u *User) GetUsers(db *sql.DB) (*[]User, error) {
	defer db.Close()
	var users []User
    rows, err := db.Query("SELECT id, name, login FROM users")
    if err != nil {
        return &users, err
    }
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Name, &u.Login)
		if err != nil {
			return &users, nil
		}
		users = append(users, *u)
	}
	return &users, nil
}


func (u *User) GetUserById(db *sql.DB, user_id int64) (error) {
	defer db.Close()
	row := db.QueryRow("SELECT id, name, login FROM users WHERE id=$1", user_id)
	err := row.Scan(&u.Id, &u.Name, &u.Login)
    return err
}


func (u *User) GetUserByLogin(db *sql.DB, login string) (error) {
	defer db.Close()
	row := db.QueryRow("SELECT id, name, login FROM users WHERE login=$1", login)
	err := row.Scan(&u.Id, &u.Name, &u.Login)
    return err
}

func (u *User) AddUser(db *sql.DB) (int64, error) {
	defer db.Close()
	u.Prepare()
	u.BeforeSave()
	result, err := db.Exec("INSERT INTO users (name, login, password) values ($1,$2,$3)", u.Name, u.Login, u.Password)
	if err != nil {
        return user_id, err
    }
	user_id, err = result.LastInsertId()
	return user_id, err
}

func (u *User) EditUser(db *sql.DB) (int64, error) {
	var err error
	defer db.Close()
	u.Prepare()
	if u.Password !=""{
		u.BeforeSave()
		_, err = db.Exec("UPDATE users SET name = $1, login = $2, password = $3 WHERE id = $4", u.Name, u.Login, u.Password, u.Id)
	}else{
		_, err = db.Exec("UPDATE users SET name = $1, login = $2 WHERE id = $3", u.Name, u.Login, u.Id)
	}
	
	return u.Id, err
}

func (u *User) DeleteImage(db *sql.DB) (int64, error) {
	defer db.Close()
	_, err := db.Exec("DELETE FROM users WHERE id = $1", u.Id)
	return u.Id, err
}
