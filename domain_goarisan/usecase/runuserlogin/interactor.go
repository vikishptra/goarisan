package runuserlogin

import (
	"context"
	"time"

	"vikishptra/domain_goarisan/model/errorenum"
)

type runUserLoginInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runUserLoginInteractor{
		outport: outputPort,
	}
}

func (r *runUserLoginInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	refreshTokenJwt, tokenJwt, userObj, err := r.outport.RunLogin(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, errorenum.DataNotFound
	}
	if tokenJwt == "" {
		return nil, errorenum.SomethingError
	}

	if err := userObj.ValidateVerifyEmail(userObj.IsActive); err != nil {
		return nil, err
	}

	res.Email = userObj.Email
	res.Name = userObj.Name
	res.Now = time.Now()
	res.Token = tokenJwt
	res.RefreshToken = refreshTokenJwt
	res.RandomString = userObj.ID.String()

	return res, nil
}
