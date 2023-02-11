package runcreatepaymentmidtrans

import (
	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TranscationCreateRequest
}

type InportResponse struct {
	Item []any `json:"item"`
}
