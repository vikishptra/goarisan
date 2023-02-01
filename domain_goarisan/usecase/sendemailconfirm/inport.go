package sendemailconfirm

import (
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Email string `form:"email"`
}

type InportResponse struct {
	Message string `json:"message"`
}
