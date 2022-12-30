package entity

import (
	"strings"
	"time"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
)

type User struct {
	ID      vo.UserID `bson:"_id" json:"id"`
	Created time.Time `bson:"created" json:"created"`
	Name    string    `bson:"name" json:"name"`
	Money   int64     `bson:"money" json:"money"`
}

type UserCreateRequest struct {
	RandomString string    `json:"id"`
	Now          time.Time `json:"time"`
	Name         string    `json:"name"`
	// Money        int64     `json:"money"`
}

func (r *User) ValidateUserCreate(req UserCreateRequest) error {

	if strings.TrimSpace(req.Name) == "" {
		return errorenum.MessageNotEmpty
	}

	return nil
}

func NewUser(req UserCreateRequest) (*User, error) {

	id, err := vo.NewUserID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj User
	obj.ID = id
	obj.Created = req.Now
	obj.Name = req.Name
	obj.Money = 0

	return &obj, nil
}

type UserUpdateRequest struct {
	ID   vo.UserID `uri:"id"`
	Name string    `json:"name"`
}

func (r *User) Update(req UserUpdateRequest) error {
	r.Name = req.Name
	if strings.TrimSpace(req.Name) == "" {
		return errorenum.MessageNotEmpty
	}
	return nil
}
