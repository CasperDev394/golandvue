package model

import "github.com/google/uuid"

type User struct {
	UID      uuid.UUID `db:"uid" json:"uid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"-"`
	Name     string    `db:"name" json:"namr"`
	ImageURl string    `db:"imsge_url" json:"imsgeUrl"`
	Website  string    `db:"webdite" json:"website"`
}
