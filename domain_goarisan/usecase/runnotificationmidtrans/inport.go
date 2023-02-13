package runnotificationmidtrans

import "vikishptra/shared/gogen"

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	OrderID       string
	TrcIDMidtrans string
}

type InportResponse struct {
	Message string `json:"message"`
}
