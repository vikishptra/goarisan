package verifyemail

import "vikishptra/shared/gogen"

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Code string `form:"code"`
	Id   string `form:"id"`
}

type InportResponse struct {
	Message string `json:"message"`
}
