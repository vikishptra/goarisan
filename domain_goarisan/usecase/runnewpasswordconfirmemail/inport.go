package runnewpasswordconfirmemail

import "vikishptra/shared/gogen"

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Code            string `json:"token"`
	Id              string `json:"id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type InportResponse struct {
	Message string `json:"message"`
}
