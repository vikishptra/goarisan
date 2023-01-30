package runlogoutuser

import (
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Token  vo.UserID `json:"token"`
	UserId []string
}

type InportResponse struct {
	Message string `json:"message"`
}
