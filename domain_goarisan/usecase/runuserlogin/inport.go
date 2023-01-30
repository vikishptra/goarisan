package runuserlogin

import (
	"time"

	"vikishptra/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Now      time.Time `json:"time"`
}

type InportResponse struct {
	Token        string    `json:"access_token"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Now          time.Time `json:"time"`
	RandomString string    `json:"id"`
	RefreshToken string    `json:"refresh_token"`
}
