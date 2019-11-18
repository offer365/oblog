package dao

import (
	"errors"

	"../model"
)

func InsertLeave(leave *model.Leave) (err error) {

	query := "insert into `leave`(username,email,content)values(?,?,?)"
	_, err = DB.Exec(query, leave.Username, leave.Email, leave.Content)
	if err != nil {
		err = errors.New("invalid parameter")
		return
	}

	return
}

func GetLeaveList() (leaveList []*model.Leave, err error) {

	query := "select id, username, email, content, create_at from `leave` order by id desc"
	err = DB.Select(&leaveList, query)
	if err != nil {
		err = errors.New("invalid parameter")
		return
	}

	return
}
