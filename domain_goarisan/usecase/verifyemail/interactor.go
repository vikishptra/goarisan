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

	// code your usecase definition here ...
	//!

	return res, nil
}
