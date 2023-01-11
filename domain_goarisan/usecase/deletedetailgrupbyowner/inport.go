package deletedetailgrupbyowner

import (
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	IDUserDetailGrup vo.DetailGrupArisanID `form:"id_detail_user"`
}

type InportResponse struct {
	Message string `json:"message"`
}
