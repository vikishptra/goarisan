package runkocokgruparisan

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	IDGrup   vo.GruparisanID `form:"id_grup"`
	IDUser   vo.UserID       `uri:"id"`
	JwtToken vo.UserID       `uri:"id"`
	entity.DetailGrupArisan
}

type InportResponse struct {
	Items []any
}
