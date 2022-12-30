package runkocokgruparisan

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	IDGrup string `uri:"id"`
	entity.DetailGrupArisan
}

type InportResponse struct {
	Items []any
}
