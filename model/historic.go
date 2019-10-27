package model

import "database/sql"

//User struct
type UserHistoric struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

//Question struct
type Question struct {
	Description string `json:"description"`
}

//Historic struct
type Historic struct {
	Id          int64        `json:"id"`
	User        UserHistoric `json:"user"`
	Question    Question     `json:"question"`
	Status      int64        `json:"status"`
	Alternative string       `json:"alternative"`
}

//GetHistoric ...
func (a *Historic) GetHistoric(db *sql.DB) error {
	err := db.QueryRow(`SELECT h.id, u.Id as userId, u.name as userName, q.description, h.status, h.alternative
					FROM historic h
					INNER JOIN users u on h.user_id = u.Id
					INNER JOIN questions q on h.question_id = q.Id
					WHERE user.Id =  ?`, a.User.Id).Scan(&a.Id, &a.User.Id, &a.User.Name, &a.Question.Description, &a.Status, &a.Alternative)
	if err != nil {
		return err
	}

	return err
}
