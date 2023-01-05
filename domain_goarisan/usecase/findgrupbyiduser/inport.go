package findgrupbyiduser

import (
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	UserID   vo.UserID `uri:"id"`
	JwtToken vo.UserID
}

type InportResponse struct {
	Item []any `json:"item"`
}
