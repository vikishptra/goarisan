package runupdateownergrup

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	IDUser  vo.UserID       `json:"id_user"`
	IDGrup  vo.GruparisanID `json:"id_grup"`
	IDOwner vo.UserID
	entity.DetailGrupArisan
}

type InportResponse struct {
	Message string `json:"message"`
}
