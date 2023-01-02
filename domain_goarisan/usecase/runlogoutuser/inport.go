package runlogoutuser

import (
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Token  string `json:"token"`
	UserId []string
}

type InportResponse struct {
	Message string
}
