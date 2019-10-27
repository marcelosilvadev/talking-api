package model

import (
	"database/sql"
	"strings"
)

//Responses struct
type Responses struct {
	Id          int64  `json: "id"`
	Alternative string `json: "alternative"`
	Description string `json: "description"`
	VideoUrl    string `json: "video_url"`
}

//Group struct
type Group struct {
	Id          int64  `json: "id"`
	Description string `json: "description"`
}

//Questions struct
type Questions struct {
	Id                 int64       `json:"id"`
	Description        string      `json:"description"`
	VideoUrl           string      `json:"video_url"`
	CorrectAlternative string      `json: "correct_alternative"`
	Points             int64       `json: "points"`
	Status             int64       `json: "status"`
	Group              Group       `json: group`
	Responses          []Responses `json:"responses"`
}

//GetQuestions ...
func (d *Questions) GetQuestions(db *sql.DB) ([]Questions, error) {
	var values []interface{}
	var where []string

	where = append(where, "status = ?")
	values = append(values, 1)

	rows, err := db.Query(`SELECT 
														q.id,
														q.description,
														q.video_url,
														q.correct_alternative,
														q.points,
														q.status,
														g.id as idGroup,
														g.description as descriptionGroup
													FROM questions q
																INNER JOIN ` + " `group` " + `g ON q.group_id = g.id WHERE 1 = 1 `)

	if err != nil {
		return nil, err
	}

	questions := []Questions{}
	defer rows.Close()
	for rows.Next() {
		var dm Questions
		if err = rows.Scan(&dm.Id, &dm.Description, &dm.VideoUrl, &dm.CorrectAlternative, &dm.Points, &dm.Status, &dm.Group.Id, &dm.Group.Description); err != nil {
			return nil, err
		}

		err := dm.GetResponses(db)
		if err != nil {
			return nil, err
		}
		questions = append(questions, dm)

	}
	return questions, nil

}

//GetResponses ...
func (d *Questions) GetResponses(db *sql.DB) error {
	var values []interface{}
	var where []string

	where = append(where, "question_id = ?")
	values = append(values, d.Id)

	rows, err := db.Query(`SELECT id,
																alternative,
																description,
																video_url																
															FROM responses
															WHERE 1 = 1 AND `+strings.Join(where, " AND "), values...)

	if err != nil {
		return err
	}

	responses := []Responses{}
	defer rows.Close()
	for rows.Next() {
		var r Responses
		if err = rows.Scan(&r.Id, &r.Alternative, &r.Description, &r.VideoUrl); err != nil {
			return err
		}

		responses = append(responses, r)

	}
	d.Responses = responses
	return nil
}
