package runuserupdate

import (
	"context"

	"vikishptra/shared/util"
)

type runUserUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runUserUpdateInteractor{
		outport: outputPort,
	}
}

func (r *runUserUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObjs, err := r.outport.FindUserByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err := userObjs.Update(req.UserUpdateRequest); err != nil {
		return nil, err
	}

	if err := r.outport.SaveUser(ctx, userObjs); err != nil {
		return nil, err
	}

	res.ID = userObjs.ID
	res.Nama = userObjs.Name
	message := []any{
		"ok success update profile",
	}
	res.Message = util.ToSliceAny(message)

	//!

	return res, nil
}
