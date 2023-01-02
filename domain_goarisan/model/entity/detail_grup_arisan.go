package entity

import (
	"math/rand"
	"strings"
	"time"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
)

type DetailGrupArisan struct {
	ID                      vo.DetailGrupArisanID `bson:"_id" json:"id"`
	ID_Detail_Grup          vo.GruparisanID       `json:"id_detail_grup" uri:"id"`
	ID_User                 vo.UserID             `json:"id_user" uri:"id"`
	StatusUserPutaranArisan int64                 `bson:"status_user_arisan"`
	Money                   int64                 `json:"money"`
	No_undian               int                   `json:"no_undian"`
	Created                 time.Time             `bson:"created" json:"created"`
}

type DetailGrupArisanCreateRequest struct {
	RandomString   string          `json:"-"`
	Now            time.Time       `json:"-"`
	ID_Detail_Grup vo.GruparisanID `json:"id_detail_grup" uri:"id"`
	ID_User        vo.UserID       `json:"id_user" uri:"id"`
	RulesMoney     int64           `json:"money"`
	StatusUser     bool            `bson:"status_user_arisan"`
	JwtToken       vo.UserID       `json:"json"`
}

type DetailGrupArisanUpdateRequest struct {
}

func JoinGrupArisan(req DetailGrupArisanCreateRequest) (*DetailGrupArisan, error) {

	id, err := vo.NewDetailGrupArisanID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var Gruparisan DetailGrupArisan
	Gruparisan.ID = id
	Gruparisan.ID_User = req.ID_User
	Gruparisan.ID_Detail_Grup = req.ID_Detail_Grup
	Gruparisan.Created = req.Now
	Gruparisan.Money = req.RulesMoney
	return &Gruparisan, nil
}
func (r *DetailGrupArisan) ValidateGrupJoin(req DetailGrupArisanCreateRequest, reqUser *User, uang int64) error {
	if reqUser.Money == 0 {
		return errorenum.MoneyMin
	} else if strings.TrimSpace(string(req.ID_Detail_Grup)) == "" || strings.TrimSpace(string(req.ID_User)) == "" {
		return errorenum.MessageNotEmpty
	} else if uang >= reqUser.Money {
		return errorenum.UserStrapped
	}
	return nil
}

func (r *DetailGrupArisan) Update(req DetailGrupArisanUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
func (r *DetailGrupArisan) SetDetailGrup(req *Gruparisan, reqRand DetailGrupArisanCreateRequest, reqUser *User) (*DetailGrupArisan, error) {
	id, err := vo.NewDetailGrupArisanID(reqRand.RandomString, req.Created)
	if err != nil {
		return nil, err
	}

	r.ID = id
	r.Created = req.Created
	r.ID_Detail_Grup = req.ID
	r.ID_User = reqUser.ID
	r.Created = req.Created
	r.Money = req.RulesMoney
	r.No_undian = rand.Intn(100)
	return r, nil

}
