package runupdategruparisanbyidowner

import (
	"context"
)

type runupdategruparisanbyidownerInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runupdategruparisanbyidownerInteractor{
		outport: outputPort,
	}
}

func (r *runupdategruparisanbyidownerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	gruparisanObj, err := r.outport.FindOneGrupByOwner(ctx, req.IDUser, req.IDGrup)
	if err != nil {
		return nil, err
	}
	if err := gruparisanObj.Update(req.GruparisanUpdateRequest); err != nil {
		return nil, err
	}
	if err := r.outport.SaveGrupArisan(ctx, gruparisanObj); err != nil {
		return nil, err
	}

	res.Message = "ok success update grup arisan"
	return res, nil
}
