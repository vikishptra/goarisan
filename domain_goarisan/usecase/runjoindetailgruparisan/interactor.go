package runjoindetailgruparisan

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
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
	_, err := r.outport.FindUserByID(ctx, req.ReqDetail.ID_User)
	if err != nil {
		return nil, err
	}
	if _, err := r.outport.FindGrupArisanById(ctx, req.ReqDetail.ID_Detail_Grup); err != nil {
		return nil, err
	}

	detailGrup, err := entity.JoinGrupArisan(req.ReqDetail)
	if err != nil {
		return nil, err
	}
	if err := detailGrup.ValidateGrupJoin(req.ReqDetail); err != nil {
		return nil, err
	}

	if err := r.outport.SaveDetailGrupArisan(ctx, detailGrup); err != nil {
		return nil, err
	}

	res.Message = "ok success"
	return res, nil
}
