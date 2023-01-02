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

	res.Message = "ok success logout"
	return res, nil
}
