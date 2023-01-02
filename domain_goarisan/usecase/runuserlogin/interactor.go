package runuserlogin

import (
	"context"

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

	// code your usecase definition here ...

	tokenJwt, userObj, err := r.outport.RunLogin(ctx, req.Name, req.Password)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, errorenum.DataNotFound
	}
	if tokenJwt == "" {
		return nil, errorenum.SomethingError
	}

	res.Name = userObj.Name
	res.Now = req.Now
	res.Password = "-"
	res.Token = tokenJwt
	res.RandomString = userObj.ID.String()
	//!

	return res, nil
}
