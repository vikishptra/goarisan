package runverifynewpasswordemail

import "vikishptra/shared/gogen"

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Code string `form:"code"`
	Id   string `form:"id"`
}

type InportResponse struct {
	Code string `json:"code"`
	Id   string `json:"id"`
}
