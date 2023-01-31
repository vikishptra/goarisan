package verifyemail

import (
	"context"
)

type verifyEmailInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &verifyEmailInteractor{
		outport: outputPort,
	}
}

func (r *verifyEmailInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	err := r.outport.RunVerifyEmail(ctx, req.Id, req.Code)
	if err != nil {
		return nil, err
	}

	res.Message = "ok success verifikasi email anda"
	return res, nil
}
