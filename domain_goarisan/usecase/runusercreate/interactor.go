package runusercreate

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
)

type runUserCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runUserCreateInteractor{
		outport: outputPort,
	}
}

func (r *runUserCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := entity.NewUser(req.UserCreateRequest)
	if err != nil {
		return nil, err
	}
	if err := todoObj.ValidateUserCreate(req.UserCreateRequest); err != nil {
		return nil, err
	}

	if err := r.outport.SaveUser(ctx, todoObj); err != nil {
		return nil, err
	}
	res.Message = "ok success"

	return res, nil
}
