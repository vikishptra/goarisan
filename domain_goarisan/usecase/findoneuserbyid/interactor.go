package findoneuserbyid

import (
	"context"
)

type findOneUserByIDInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &findOneUserByIDInteractor{
		outport: outputPort,
	}
}

func (r *findOneUserByIDInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObjByID, err := r.outport.FindUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if err := userObjByID.ValidateTokenUser(req.UserID, req.JwtToken); err != nil {
		return nil, err
	}
	res.Email = userObjByID.Email
	res.IsActive = userObjByID.IsActive
	res.Name = userObjByID.Name

	return res, nil
}
