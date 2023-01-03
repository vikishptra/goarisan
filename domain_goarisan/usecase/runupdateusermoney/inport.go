package runupdateusermoney

import (
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	UserID   vo.UserID `uri:"id"`
	Money    int64     `json:"money"`
	JwtToken vo.UserID
}

type InportResponse struct {
	Message string `json:"message"`
}
