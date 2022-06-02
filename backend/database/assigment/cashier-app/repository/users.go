package repository

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {

	sqlStmt := `SELECT * FROM users WHERE id = ?`
	var user User
	row := u.db.QueryRow(sqlStmt, id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
	if err != nil {
		return User{}, err
	}
	return user, nil
	//return User{}, nil // TODO: replace this
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	sqlStmt := `SELECT * FROM users`
	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
	//return []User{}, nil // TODO: replace this
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	sqlStmt := `SELECT * FROM users WHERE username = ? AND password = ?`

	row := u.db.QueryRow(sqlStmt, username, password)
	user := User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
	if err != nil {
		return nil, errors.New("Login Failed")
	}
	if user.Loggedin == true {
		return nil, errors.New("User already logged in")
	}
	sqlStmt = `UPDATE users SET loggedin = ? WHERE username = ?`
	_, err = u.db.Exec(sqlStmt, true, username)
	if err != nil {
		return nil, err
	}
	return &user.Username, nil

}

func (u *UserRepository) InsertUser(username string, password string, role string, loggedin bool) error {

	sqlStmt := `INSERT INTO users (username, password, role, loggedin) VALUES (?, ?, ?, ?)`
	_, err := u.db.Exec(sqlStmt, username, password, role, loggedin)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {

	sqlStmt := `SELECT role FROM users WHERE username = ?`
	var role string
	err := u.db.QueryRow(sqlStmt, username).Scan(&role)
	if err != nil {
		return nil, err
	}

	return &role, nil // TODO: replace this

}
