package findoneuserbyid

import (
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	UserID   vo.UserID
	JwtToken vo.UserID
}

type InportResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
