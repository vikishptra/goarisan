package entity

import (
	"strings"
	"time"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
)

type Gruparisan struct {
	ID          vo.GruparisanID `bson:"_id" json:"id"`
	Created     time.Time       `bson:"created" json:"created"`
	NamaGrup    string          `json:"nama_grup"`
	ID_Owner    vo.UserID       `uri:"id" json:"id_owner"`
	JumlahUsers int64           `json:"jumlah_users"`
	RulesMoney  int64           `json:"rules_money"`
	ResultMoney int64           `json:"result_money"`
}

type GruparisanCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	ID_Owner     vo.UserID `uri:"id" json:"id_owner"`
	NamaGrup     string    `json:"nama_grup"`
	JumlahUsers  int64     `json:"jumlah_users"`
	RulesMoney   int64     `json:"rules_money"`
	ResultMoney  int64     `json:"result_money"`
}

func (r *Gruparisan) ValidateGrupCreate(req GruparisanCreateRequest, reqUser *User) error {

	if reqUser.Money == 0 {
		return errorenum.MoneyMin
	} else if strings.TrimSpace(req.NamaGrup) == "" {
		return errorenum.MessageNotEmpty
	} else if req.RulesMoney >= reqUser.Money {
		return errorenum.UserStrapped
	}

	return nil
}

func NewGruparisan(req GruparisanCreateRequest) (*Gruparisan, error) {

	id, err := vo.NewGruparisanID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Gruparisan
	obj.ID = id
	obj.NamaGrup = req.NamaGrup
	obj.Created = req.Now
	obj.ID_Owner = req.ID_Owner
	obj.JumlahUsers = req.JumlahUsers
	obj.ResultMoney = req.ResultMoney
	obj.RulesMoney = req.RulesMoney

	return &obj, nil
}

type GruparisanUpdateRequest struct {
	// add field to update here ...
}

func (r *Gruparisan) Update(req GruparisanUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}

func (r *Gruparisan) SetIdUser(req *User) error {
	r.ID_Owner = req.ID

	return nil
}
func (g *Gruparisan) UpdateMoneyUserGrup(reqRules int64, r *User) error {

	if int64(r.Money) == 0 {
		return errorenum.MoneyMin
	}
	r.Money = r.Money - reqRules

	return nil

}
