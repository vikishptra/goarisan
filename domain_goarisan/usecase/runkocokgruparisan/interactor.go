package runkocokgruparisan

import (
	"context"

	"vikishptra/shared/util"
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

	detailGrupObj, err := r.outport.FindUndianArisanUser(ctx, req.IDGrup)
	if err != nil {
		return nil, err
	}
	res.Items = util.ToSliceAny(detailGrupObj)

	return res, nil
}
