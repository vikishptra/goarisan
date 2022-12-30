package rungruparisancreate

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
)

type runGrupArisanCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runGrupArisanCreateInteractor{
		outport: outputPort,
	}
}

func (r *runGrupArisanCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObjs, err := r.outport.FindUserByID(ctx, req.ID_Owner)
	if err != nil {
		return nil, err
	}

	grupObj, err := entity.NewGruparisan(req.GruparisanCreateRequest)
	if err != nil {
		return nil, err
	}
	grupObj.SetIdUser(userObjs)
	if err := grupObj.ValidateGrupCreate(req.GruparisanCreateRequest, userObjs); err != nil {
		return nil, err
	}
	var RulesMoney = req.GruparisanCreateRequest.RulesMoney

	grupObj.UpdateMoneyUserGrup(RulesMoney, userObjs)

	if err := r.outport.SaveGrupArisan(ctx, grupObj); err != nil {
		return nil, err
	}

	if err := r.outport.SaveUser(ctx, userObjs); err != nil {
		return nil, err
	}
	detailGrup, err := req.Detail.SetDetailGrup(grupObj, req.DetailReq, userObjs)
	if err != nil {
		return nil, err
	}
	if err := r.outport.SaveDetailGrupArisan(ctx, detailGrup); err != nil {
		return nil, err
	}

	res.Message = "ok success"

	return res, nil
}
