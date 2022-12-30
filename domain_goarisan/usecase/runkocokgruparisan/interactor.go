package runkocokgruparisan

import (
	"context"
)

type runKocokGrupArisanInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runKocokGrupArisanInteractor{
		outport: outputPort,
	}
}

func (r *runKocokGrupArisanInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	return res, nil
}
