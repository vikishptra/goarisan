package runupdateusermoney

import (
	"context"
)

type runupdateusermoneyInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runupdateusermoneyInteractor{
		outport: outputPort,
	}
}

func (r *runupdateusermoneyInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObjByID, err := r.outport.FindUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if err := userObjByID.ValidateTokenUser(req.UserID, req.JwtToken); err != nil {
		return nil, err
	}
	if err := userObjByID.UpdateMoney(req.Money); err != nil {
		return nil, err
	}
	if err := r.outport.SaveUser(ctx, userObjByID); err != nil {
		return nil, err
	}
	res.Message = "ok success update money"
	return res, nil
}
