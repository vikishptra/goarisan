package entity

import (
	"strings"
	"time"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
)

type DetailGrupArisan struct {
	ID             vo.DetailGrupArisanID `bson:"_id" json:"id"`
	ID_Detail_Grup vo.GruparisanID       `json:"id_detail_grup" uri:"id"`
	ID_User        vo.UserID             `json:"id_user" uri:"id"`
	StatusUser     bool                  `bson:"status_user_arisan"`
	Created        time.Time             `bson:"created" json:"created"`
}

type DetailGrupArisanCreateRequest struct {
	RandomString   string          `json:"-"`
	Now            time.Time       `json:"-"`
	ID_Detail_Grup vo.GruparisanID `json:"id_detail_grup" uri:"id"`
	ID_User        vo.UserID       `json:"id_user" uri:"id"`
	StatusUser     bool            `bson:"status_user_arisan"`
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
	return &Gruparisan, nil
}
func (r *DetailGrupArisan) ValidateGrupJoin(req DetailGrupArisanCreateRequest) error {

	if strings.TrimSpace(string(req.ID_Detail_Grup)) == "" || strings.TrimSpace(string(req.ID_User)) == "" {
		return errorenum.MessageNotEmpty
	}
	return nil
}

func (r *DetailGrupArisan) Update(req DetailGrupArisanUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
func (r *DetailGrupArisan) SetDetailGrup(req *Gruparisan, reqRand DetailGrupArisanCreateRequest) (*DetailGrupArisan, error) {
	id, err := vo.NewDetailGrupArisanID(reqRand.RandomString, req.Created)
	if err != nil {
		return nil, err
	}
	r.ID = id
	r.Created = req.Created
	r.ID_Detail_Grup = req.ID
	r.ID_User = req.ID_Owner
	r.Created = req.Created

	return r, nil

}
