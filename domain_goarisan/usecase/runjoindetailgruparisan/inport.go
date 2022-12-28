package runjoindetailgruparisan

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	ReqDetail entity.DetailGrupArisanCreateRequest
}

type InportResponse struct {
	Message string
}
