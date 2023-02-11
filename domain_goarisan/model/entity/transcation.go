package entity

import (
	"time"

	"vikishptra/domain_goarisan/model/vo"
)

type Transcation struct {
	ID               vo.TransaksiID `bson:"_id" json:"id"`
	IDUser           vo.UserID      `bson:"id_user" json:"id_user"`
	MoneyUser        int64          `bson:"money_user" json:"money_user"`
	Bank             string         `json:"bank"`
	ResponseMidtrans string         `bson:"response_midtrans" json:"response_midtrans"`
	Updated          time.Time      `bson:"updated" json:"updated"`
	Created          time.Time      `bson:"created" json:"created"`
}

type TranscationCreateRequest struct {
	RandomString string    `json:"-"`
	Bank         string    `json:"bank"`
	IDuser       vo.UserID `json:"id_user"`
	MoneyUser    int64     `json:"money_user"`
	Now          time.Time `json:"-"`
}

func NewTranscation(req TranscationCreateRequest) (*Transcation, error) {

	id, err := vo.NewTransaksiID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Transcation
	obj.ID = id
	obj.Bank = req.Bank
	obj.MoneyUser = req.MoneyUser
	obj.IDUser = req.IDuser
	obj.Updated = req.Now
	obj.Created = req.Now

	return &obj, nil
}

type TranscationUpdateRequest struct {
	// add field to update here ...
}

func (r *Transcation) Update(req TranscationUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
