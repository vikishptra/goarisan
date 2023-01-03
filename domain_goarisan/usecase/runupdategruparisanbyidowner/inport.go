package runupdategruparisanbyidowner

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.GruparisanUpdateRequest
}

type InportResponse struct {
	Message string `json:"message"`
}
