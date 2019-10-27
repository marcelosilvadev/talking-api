package model

import (
	"database/sql"
)

//User struct
type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Photo    string `json: "photo"`
	Points   string `json: "points"`
	Password string `json: "password"`
}

//InsertUser ...
func (u *User) InsertUser(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO USERS (NAME, EMAIL, PHOTO, POINTS)
								VALUES
								(?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(u.Name, u.Email, u.Photo, u.Points)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	u.Id = id
	return nil
}

//GetUser ...
func (u *User) GetUser(db *sql.DB) error {
	err := db.QueryRow(`select id, name, email, photo, points, password
					from users
					where email =  ?`, u.Email).Scan(&u.Id, &u.Name, &u.Email, &u.Photo, &u.Points, &u.Password)
	if err != nil {
		return err
	}

	return err
}

//Login ...
func (u *User) Login(db *sql.DB) error {
	err := db.QueryRow(`select id, name, email, photo, points, password
					from users
					where email =  ? and password = ?`, u.Email, u.Password).Scan(&u.Id, &u.Name, &u.Email, &u.Photo, &u.Points, &u.Password)
	if err != nil {
		return err
	}

	return err
}

//GetRanking ...
func (a *User) GetRanking(db *sql.DB) ([]User, error) {

	rows, err := db.Query(`select id, name, points
													from users
													order by points desc `)

	if err != nil {
		return nil, err
	}

	ranking := []User{}
	defer rows.Close()
	for rows.Next() {
		var as User
		if err = rows.Scan(&as.Id, &as.Name, &as.Points); err != nil {
			return nil, err
		}
		ranking = append(ranking, as)
	}
	return ranking, nil
}

//UpdatePoints ...
func (a *User) UpdatePoints(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE users SET points = points + ? WHERE id = ?;`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(a.Points, a.Id)

	return err
}
