package runjoindetailgruparisan

import (
	"context"
)

type runJoinDetailGrupArisanInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runJoinDetailGrupArisanInteractor{
		outport: outputPort,
	}
}

func (r *runJoinDetailGrupArisanInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	return res, nil
}
