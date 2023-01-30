package refreshtokenjwt

import (
	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	IDUser string
	Token  string
}

type InportResponse struct {
	AccessToken string `json:"access_token"`
}
