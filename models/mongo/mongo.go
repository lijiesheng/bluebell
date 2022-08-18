package mongo

import "time"

type CcSheets struct {
	Id string `bson:"_id,omitempty" json:"id"`
	AdviserNo string `bson:"adviser_no" json:"adviser_no"`
	CallNo string `bson:"call_no" json:"call_no"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	Mobile string `bson:"mobile" json:"mobile"`
	Remark string `bson:"remark" json:"remark"`
	Sdate string `bson:"sdate" json:"sdate"`
	Status int64 `bson:"status" json:"status"`
	TlAccount string `bson:"tl_account" json:"tl_account"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	V int64 	`bson:"__v" json:"__v"`
}

