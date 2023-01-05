package findgruparisanbyidowner

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Grup     entity.Gruparisan
	JwtToken vo.UserID
}

type InportResponse struct {
	Item []any `json:"item"`
}
