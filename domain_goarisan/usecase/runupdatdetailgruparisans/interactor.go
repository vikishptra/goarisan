package runupdatdetailgruparisans

import (
	"context"
)

type runupdatdetailgruparisansInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runupdatdetailgruparisansInteractor{
		outport: outputPort,
	}
}

func (r *runupdatdetailgruparisansInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	detailObj, err := r.outport.Findoneuserdetailgruparisans(ctx, req.ID_Detail_Grup, req.ID_User)
	if err != nil {
		return nil, err
	}

	if err := detailObj.ValidateTokenUser(req.ID_User, req.JwtToken); err != nil {
		return nil, err
	}

	grupObj, err := r.outport.FindGrupArisanyIdGrup(ctx, req.ID_Detail_Grup)
	if err != nil {
		return nil, err
	}
	if err := detailObj.UpdateDetailGrupUser(req.DetailGrupArisanCreateRequest, grupObj.RulesMoney); err != nil {
		return nil, err
	}

	userObj, err := r.outport.FindUserByID(ctx, req.ID_User)
	if err != nil {
		return nil, err
	}

	if err := grupObj.UpdateMoneyUserGrup(grupObj.RulesMoney, userObj); err != nil {
		return nil, err
	}

	if err := r.outport.SaveDetailGrupArisan(ctx, detailObj); err != nil {
		return nil, err
	}
	if err := r.outport.SaveUser(ctx, userObj); err != nil {
		return nil, err
	}

	res.Message = "ok success update"

	return res, nil
}
