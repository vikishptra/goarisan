package runverifynewpasswordemail

import (
	"context"
)

type runverifynewpasswordemailInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runverifynewpasswordemailInteractor{
		outport: outputPort,
	}
}

func (r *runverifynewpasswordemailInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	if err := r.outport.RunVerifyNewPasswordEmail(ctx, req.Id, req.Code); err != nil {
		return nil, err
	}
	res.Code = req.Code
	res.Id = req.Id
	return res, nil
}
