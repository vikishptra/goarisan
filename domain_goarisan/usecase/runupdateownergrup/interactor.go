package runupdateownergrup

import (
	"context"
)

type runUpdateOwnerGrupInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runUpdateOwnerGrupInteractor{
		outport: outputPort,
	}
}

func (r *runUpdateOwnerGrupInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	detailGrupObj, err := r.outport.RunUpdateOwnerGrup(ctx, req.IDUser, req.IDGrup, req.IDOwner)
	if err != nil {
		return nil, err
	}

	if err := req.ValidateUserSame(req.IDOwner, req.IDUser); err != nil {
		return nil, err
	}

	if err := detailGrupObj.UpdateOwnerGrup(req.IDUser); err != nil {
		return nil, err
	}
	if err := r.outport.SaveGrupArisan(ctx, detailGrupObj); err != nil {
		return nil, err
	}

	res.Message = "Owner Berhasil Di Update!"

	return res, nil
}
