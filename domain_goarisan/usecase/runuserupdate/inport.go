package runuserupdate

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.UserUpdateRequest
}

type InportResponse struct {
	Items *entity.User
}
