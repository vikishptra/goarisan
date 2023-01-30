package findoneuserbyid

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/shared/util"
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
	var users []entity.User

	res := &InportResponse{}

	userObjByID, err := r.outport.FindUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if err := userObjByID.ValidateTokenUser(req.UserID, req.JwtToken); err != nil {
		return nil, err
	}
	userObjByID.Password = "-"
	users = append(users, *userObjByID)

	res.Item = util.ToSliceAny(users)

	return res, nil
}
