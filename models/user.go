package models

import "time"

type User struct {
	Id int64
	User_id int64
	Username string
	Password string
	Email string
	Gender int8
	Create_time time.Time
	Update_time time.Time
}