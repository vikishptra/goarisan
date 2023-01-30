package runlogoutuser

import (
	"context"
)

type runLogoutUserInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runLogoutUserInteractor{
		outport: outputPort,
	}
}

func (r *runLogoutUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	if err := r.outport.RunLogout(ctx, req.Token); err != nil {
		return nil, err
	}

	res.Message = "ok success logout"
	return res, nil
}
