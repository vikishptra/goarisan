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
	userObjs, err := r.outport.FindUserByID(ctx, req.ReqDetail.ID_User)
	if err != nil {
		return nil, err
	}
	grupObjs, err := r.outport.FindGrupArisanAndUserById(ctx, req.ReqDetail.ID_Detail_Grup, userObjs.ID)
	if err != nil {
		return nil, err
	}
	var RulesMoney = grupObjs.RulesMoney

	detailGrup, err := entity.JoinGrupArisan(req.ReqDetail)
	if err != nil {
		return nil, err
	}

	if err := detailGrup.ValidateGrupJoin(req.ReqDetail, userObjs, RulesMoney); err != nil {
		return nil, err
	}
	grupObjs.UpdateMoneyUserGrup(RulesMoney, userObjs)

	if err := r.outport.SaveDetailGrupArisan(ctx, detailGrup); err != nil {
		return nil, err
	}
	if err := r.outport.SaveGrupArisan(ctx, grupObjs); err != nil {
		return nil, err
	}
	if err := r.outport.SaveUser(ctx, userObjs); err != nil {
		return nil, err
	}

	res.Message = "ok success"
	return res, nil
}
