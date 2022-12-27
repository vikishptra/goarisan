package rungruparisancreate

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.GruparisanCreateRequest
	DetailReq entity.DetailGrupArisanCreateRequest
	Detail    entity.DetailGrupArisan
}

type InportResponse struct {
	Message string `json:"message"`
}
