package runuserupdate

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.UserUpdateRequest
}

type InportResponse struct {
	ID      vo.UserID `json:"id"`
	Nama    string    `json:"nama"`
	Message []any     `json:"message"`
}
